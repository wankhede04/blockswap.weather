package weatherservice

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	"github.com/gin-gonic/gin"
	"github.com/wankhede04/blockswap.weather/weather-srv/db"
)

type WeatherReport struct {
	Address   string `json:"address"`
	Report    string `json:"report"`
	Signature string `json:"signature"`
}

// AuthenticateMiddleware checks if user is registered on contract and verify data provided by user
func (s *WeatherService) AuthenticateMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		payload := WeatherReport{}
		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
			c.Abort()
			return
		}

		address := payload.Address

		if err := VerifyOrderSignature(payload, s.worker.GetChainID(), s.worker.GetRegistrationContract().String()); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Error in verification"})
			c.Abort()
			return
		}

		// Acquire a database connection
		database, err := s.getDBConnection()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		// Release the database connection
		defer s.releaseDBConnection(database)
		var membership db.Membership
		if err := database.Where("address = ?", address).First(&membership).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		if membership.Status != string(db.Registered) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// s.releaseDBConnection(database)

		c.Set("membership", membership)
		c.Set("report", payload.Report)

		c.Next()
	}
}

func VerifyOrderSignature(weatherReport WeatherReport, chainID int64, peripheryContract string) error {
	hash, err := EncodeOrderStruct(weatherReport, chainID, peripheryContract)
	if err != nil {
		return err
	}

	sign, err := hexutil.Decode(weatherReport.Signature)
	if err != nil {
		return err
	}

	if sign[64] >= 27 {
		sign[64] -= 27 // Transform V from 0/1 to 27/28 according to the yellow paper
	}

	sigPubKeyBytes, err := crypto.Ecrecover(hash, sign)
	if err != nil {
		return err
	}

	var buf []byte

	hash0 := crypto.Keccak256Hash(sigPubKeyBytes[1:])
	buf = hash0.Bytes()
	publicAddress := hexutil.Encode(buf[12:])

	hex := strings.ToLower(publicAddress)[2:]
	checkSummedAddress := "0x"
	hashCs := crypto.Keccak256Hash([]byte(hex))

	for i, b := range hex {
		c := string(b)
		if b < '0' || b > '9' {
			if hashCs[i/2]&byte(128-i%2*120) != 0 {
				c = string(b - 32)
			}
		}
		checkSummedAddress += c
	}

	result := strings.EqualFold(checkSummedAddress, weatherReport.Address)
	if !result {
		return fmt.Errorf("signer != trader")
	}
	return nil
}

// EncodeOrderStruct encodes order struct in bytes
func EncodeOrderStruct(report WeatherReport, chainID int64, positioningContract string) ([]byte, error) {
	typeddata := apitypes.TypedData{
		Types: apitypes.Types{
			"EIP712Domain": []apitypes.Type{
				{Name: "name", Type: "string"},
				{Name: "version", Type: "string"},
				{Name: "chainId", Type: "uint256"},
				{Name: "verifyingContract", Type: "address"},
			},
			"WeatherReport": []apitypes.Type{
				{Name: "address", Type: "string"},
				{Name: "report", Type: "string"},
			},
		},
		PrimaryType: "WeatherReport",
		Domain: apitypes.TypedDataDomain{
			Name:              "WeatherReport",
			Version:           "1",
			ChainId:           math.NewHexOrDecimal256(chainID),
			VerifyingContract: positioningContract,
		},
		Message: apitypes.TypedDataMessage{
			"address": report.Address,
			"report":  report.Report,
		},
	}

	rawData, _, err := apitypes.TypedDataAndHash(typeddata)

	if err != nil {
		return nil, err
	}

	return rawData, nil
}
