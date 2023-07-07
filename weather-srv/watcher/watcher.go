package watcher

import (
	"context"
	"sync"
	"time"

	"github.com/wankhede04/blockswap.weather/weather-srv/db"
	"github.com/wankhede04/blockswap.weather/weather-srv/worker"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type WatcherSRV struct {
	Logs     chan types.Log
	Sub      ethereum.Subscription
	Logger   *logrus.Logger
	DataBase *db.PostgresDataBase
	Worker   *worker.Worker
	ctx      context.Context
	cancelFn context.CancelFunc
	dbPool   *db.ConnectionPool // Custom connection pool
	dbMutex  sync.Mutex         // Mutex for database connection synchronization
}

// NewWatcherSRV creates a new WatcherSRV instance
func NewWatcherSRV(database *db.PostgresDataBase, logger *logrus.Logger, wrkr *worker.Worker) (*WatcherSRV, error) {
	logs := make(chan types.Log)

	subs, err := wrkr.SubscribeToLogs(logs)
	if err != nil {
		return nil, err
	}

	ctx, cancelFn := context.WithCancel(context.Background())

	// Create a connection pool with a maximum number of connections
	dbPool, err := db.NewConnectionPool(database.DB, 10) // Adjust the maximum number of connections as per your requirement
	if err != nil {
		cancelFn() // Call cancelFn in case of error
		return nil, err
	}

	return &WatcherSRV{
		Logs:     logs,
		Sub:      subs,
		Logger:   logger,
		DataBase: database,
		Worker:   wrkr,
		ctx:      ctx,
		cancelFn: cancelFn,
		dbPool:   dbPool,
		dbMutex:  sync.Mutex{},
	}, nil
}

// Run starts the WatcherSRV and begins processing event logs
func (w *WatcherSRV) Run() {
	go w.processEventLogs()
}

// getDBConnection acquires a database connection from the pool
func (w *WatcherSRV) getDBConnection() (*gorm.DB, error) {
	w.dbMutex.Lock()
	defer w.dbMutex.Unlock()

	db, err := w.dbPool.AcquireConnection()
	if err != nil {
		return nil, err
	}

	return db, nil
}

// releaseDBConnection releases a database connection back to the pool
func (w *WatcherSRV) releaseDBConnection(db *gorm.DB) {
	w.dbMutex.Lock()
	defer w.dbMutex.Unlock()

	w.dbPool.ReleaseConnection(db)
}

// processEventLogs continuously listens for event logs and handles them
func (w *WatcherSRV) processEventLogs() {
	for {
		select {
		case err := <-w.Sub.Err():
			w.Logger.Errorf("Error received in event subscription: %v", err)
			w.renewSubscription()
		case vLog := <-w.Logs:
			err := w.handleEventLog(vLog)
			if err != nil {
				w.Logger.Errorf("Error processing event log: %v", err)
			}
		case <-w.ctx.Done():
			w.Logger.Info("Watcher service has stopped")
			return
		}
	}
}

// renewSubscription renews the event subscription
func (w *WatcherSRV) renewSubscription() {
	for {
		subs, err := w.Worker.SubscribeToLogs(w.Logs)
		if err != nil {
			w.Logger.Errorf("Failed to renew event subscription: %v", err)
			time.Sleep(10 * time.Second)
		} else {
			w.Sub = subs
			w.Logger.Info("Event subscription renewed successfully")
			return
		}
	}
}

// handleEventLog handles an individual event log
func (w *WatcherSRV) handleEventLog(vLog types.Log) error {
	var tLog db.EventLog

	tLog.BlockHeight = vLog.BlockNumber
	tLog.TransactionHash = vLog.TxHash.Hex()

	event, eventType, err := worker.ParseEvent(&vLog)
	if err != nil {
		return err
	}
	database, err := w.getDBConnection()
	if err != nil {
		w.Logger.Errorf("Error: unable to  create connection pool %v\n", err)
	}
	defer w.releaseDBConnection(database) // Ensure the connection is released
	switch eventType {
	case "ParticipantRegistered":
		ParticipantRegistered := event.(worker.RegistrationParticipantRegistered)
		tLog.Address = ParticipantRegistered.Participant.Hex()

		membership, err := db.FindMemberShip(database, tLog.Address)
		if err != nil {
			membership := db.Membership{
				Address: tLog.Address,
				Status:  string(db.Registered),
			}
			err := db.CreateMembership(database, &membership)
			if err != nil {
				w.Logger.Errorf("Error: unable to  create %v\n", err)
			}
		} else {
			if err := db.UpdateMemberShipStatus(database, membership.Address, db.Registered); err != nil {
				w.Logger.Errorf("Error: unable to update DB %v\n", err)
			}
		}
		w.Logger.Infof("Found ParticipantRegistered event and updated membership status successfully with member %s", tLog.Address)

	case "ParticipantResigned":
		ParticipantResigned := event.(worker.RegistrationParticipantResigned)
		tLog.Address = ParticipantResigned.Participant.Hex()

		membership, err := db.FindMemberShip(database, tLog.Address)
		if err != nil {
			membership := db.Membership{
				Address: tLog.Address,
				Status:  string(db.Resigned),
			}
			err := db.CreateMembership(database, &membership)
			if err != nil {
				w.Logger.Errorf("Error: unable to  create %v\n", err)
			}
		} else {
			if err := db.UpdateMemberShipStatus(database, membership.Address, db.Resigned); err != nil {
				w.Logger.Errorf("Error: unable to update DB %v\n", err)
			}
		}

		w.Logger.Infof("Found ParticipantResigned event and updated membership status successfully with member %s", tLog.Address)
	}

	if err := db.CreateEventLog(database, &tLog); err != nil {
		w.Logger.Errorf("Error creating event log: %v", err)
	}

	return nil
}
