package db

import (
	"errors"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ConnectionPool represents a pool of database connections
type ConnectionPool struct {
	maxSize      int
	pool         chan *gorm.DB
	connectionMu sync.RWMutex
}

// NewConnectionPool creates a new connection pool with the specified maximum size
func NewConnectionPool(db *gorm.DB, maxSize int) (*ConnectionPool, error) {
	pool := make(chan *gorm.DB, maxSize)

	for i := 0; i < maxSize; i++ {
		// Clone the existing *gorm.DB connection
		clonedDB, err := db.DB()
		if err != nil {
			// If any error occurs while creating connections, release all acquired connections and return the error
			for j := 0; j < i; j++ {
				releasedDB := <-pool
				sqlDB, err := releasedDB.DB()
				if err == nil {
					sqlDB.Close()
				}
			}
			return nil, err
		}

		// Create a new *gorm.DB instance using the cloned *sql.DB connection
		clonedGormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: clonedDB}), &gorm.Config{})
		if err != nil {
			// If any error occurs while creating connections, release all acquired connections and return the error
			for j := 0; j < i; j++ {
				releasedDB := <-pool
				sqlDB, err := releasedDB.DB()
				if err == nil {
					sqlDB.Close()
				}
			}
			return nil, err
		}

		pool <- clonedGormDB
	}

	return &ConnectionPool{
		maxSize:      maxSize,
		pool:         pool,
		connectionMu: sync.RWMutex{},
	}, nil
}

// AcquireConnection acquires a database connection from the pool
func (cp *ConnectionPool) AcquireConnection() (*gorm.DB, error) {
	select {
	case db := <-cp.pool:
		return db, nil
	default:
		return nil, errors.New("connection pool is empty")
	}
}

// ReleaseConnection releases a database connection back to the pool
func (cp *ConnectionPool) ReleaseConnection(db *gorm.DB) {
	cp.pool <- db
}

// CloseConnections closes all the connections in the pool
func (cp *ConnectionPool) CloseConnections() {
	cp.connectionMu.Lock()
	defer cp.connectionMu.Unlock()

	// Check if the channel is already closed
	if cp.isPoolClosed() {
		return
	}
	close(cp.pool)
	for db := range cp.pool {
		sqlDB, err := db.DB()
		if err == nil {
			sqlDB.Close()
		}
	}
}

// isPoolClosed checks if the pool channel is closed
func (cp *ConnectionPool) isPoolClosed() bool {
	select {
	case <-cp.pool:
		return true
	default:
		return false
	}
}
