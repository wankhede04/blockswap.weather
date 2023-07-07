package worker

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
	"github.com/wankhede04/blockswap.weather/weather-srv/db"
)

// WorkerConfig ...
type WorkerConfig struct {
	ChainName            string         `json:"chain_name"`
	Provider             string         `json:"provider"`
	RegistrationContract common.Address `json:"registration_contract"`
	StartBlockHeight     *big.Int       `json:"from_block"`
}

// Worker creates an instance and store its information
type Worker struct {
	provider             string
	ChainName            string
	chainID              int64
	Logger               *logrus.Entry // Logger
	config               WorkerConfig
	client               *ethclient.Client
	registrationContract common.Address
	DB                   *db.PostgresDataBase
	Threshold            int64
}

// RegistrationParticipantRegistered represents a OrdersFilled event raised by the MatchingEngineAbi contract.
type RegistrationParticipantRegistered struct {
	Participant common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// RegistrationParticipantResigned represents a ParticipantResigned event raised by the Registration contract.
type RegistrationParticipantResigned struct {
	Participant common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// NewWorker: initialises worker (used for tx on any chain)
func NewWorker(Logger *logrus.Logger, cfg WorkerConfig, db *db.PostgresDataBase) *Worker {
	client, err := ethclient.Dial(cfg.Provider)
	if err != nil {
		panic(fmt.Sprintf("rpc error for %s : %s", cfg.ChainName, err.Error()))
	}

	chainid, err := client.ChainID(context.Background())
	if err != nil {
		panic("rpc not returning chain id")
	}

	return &Worker{
		ChainName:            cfg.ChainName,
		chainID:              chainid.Int64(),
		Logger:               Logger.WithField("worker", cfg.ChainName),
		provider:             cfg.Provider,
		config:               cfg,
		client:               client,
		registrationContract: cfg.RegistrationContract,
		DB:                   db,
	}
}

func (w *Worker) GetChainID() int64 {
	return w.chainID
}

func (w *Worker) GetRegistrationContract() common.Address {
	return w.registrationContract
}

// GetLatestBlock returns latest block
func (w *Worker) GetLatestBlock() (*big.Int, error) {
	latestBlock, err := w.client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return big.NewInt(0), err
	}
	return latestBlock.Number, nil
}

func (w *Worker) SubscribeToLogs(logs chan types.Log) (ethereum.Subscription, error) {
	var StartBlockHeight *big.Int
	lastTxnLog, err := w.DB.FindLastEventLog(w.ChainName)
	if err != nil {
		StartBlockHeight = w.config.StartBlockHeight
		if StartBlockHeight.Cmp(big.NewInt(0)) == 0 {
			StartBlockHeight, err = w.GetLatestBlock()
			if err != nil {
				return nil, fmt.Errorf("SubscribeToLogs:%w", err)
			}
		}
	} else {
		StartBlockHeight = big.NewInt(int64(lastTxnLog.BlockHeight))
	}

	// TODO : recheck this
	query := ethereum.FilterQuery{
		Addresses: []common.Address{w.config.RegistrationContract}, FromBlock: StartBlockHeight,
	}

	sub, err := w.client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		return nil, fmt.Errorf("SubscribeToLogs:%w", err)
	}

	return sub, nil
}

// ParseEvent parses an individual event log into interface
func ParseEvent(log *types.Log) (interface{}, string, error) {

	participantRegisteredSig := "ParticipantRegistered(address)"

	participantRegisteredSigHash := crypto.Keccak256Hash([]byte(participantRegisteredSig))

	participantResignedSig := "ParticipantResigned(address)"
	participantResignedSigHash := crypto.Keccak256Hash([]byte(participantResignedSig))

	switch log.Topics[0].Hex() {
	case participantRegisteredSigHash.Hex():
		participantRegistered := RegistrationParticipantRegistered{
			Participant: common.BytesToAddress(log.Topics[1].Bytes()),
		}
		return participantRegistered, "ParticipantRegistered", nil
	case participantResignedSigHash.Hex():
		ParticipantResigned := RegistrationParticipantResigned{
			Participant: common.BytesToAddress(log.Topics[1].Bytes()),
		}
		return ParticipantResigned, "ParticipantResigned", nil
	}
	return RegistrationParticipantResigned{}, "", nil

}
