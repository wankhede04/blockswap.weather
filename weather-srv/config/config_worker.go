package config

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

// WorkerConfig worker configuration struct
type WorkerConfig struct {
	ChainName            string         `json:"chain_name"`
	Provider             string         `json:"provider"`
	RegistrationContract common.Address `json:"registration_contract"`
	GasLimit             int64          `json:"gas_limit"`
	WorkerAddr           common.Address `json:"worker_addr"`
	PrivateKey           string         `json:"private_key"`
	StartBlockHeight     *big.Int       `json:"from_block"`
}

// readWorkerConfig reads ethereum chain worker params from config.json
func (v *viperConfig) readWorkerConfig(chain string) WorkerConfig {
	return WorkerConfig{
		ChainName:            strings.ToUpper(chain),
		WorkerAddr:           common.HexToAddress(v.GetString(fmt.Sprintf("workers.%s.worker_addr", chain))),
		PrivateKey:           v.GetString(fmt.Sprintf("workers.%s.private_key", chain)),
		RegistrationContract: common.HexToAddress(v.GetString(fmt.Sprintf("workers.%s.registration_contract", chain))),
		Provider:             v.GetString(fmt.Sprintf("workers.%s.provider", chain)),
		GasLimit:             v.GetInt64(fmt.Sprintf("workers.%s.gas_limit", chain)),
		StartBlockHeight:     big.NewInt(v.GetInt64(fmt.Sprintf("workers.%s.start_block_height", chain))),
	}
}

func (v *viperConfig) ReadWorkersConfig() WorkerConfig {
	return v.readWorkerConfig("ARB")
}
