package db

import (
	"time"

	"gorm.io/gorm"
)

// CreateEventLog creates a event log in DB
func CreateEventLog(DB *gorm.DB, tLog *EventLog) error {
	tLog.CreatedAt = time.Now().Unix()

	if err := DB.Create(&tLog); err.Error != nil {
		return err.Error
	}
	return nil
}

// FindLatestTxnLog returns last event log stored in DB
func (db *PostgresDataBase) FindLastEventLog(chain string) (*EventLog, error) {
	var lastEvent EventLog
	if result := db.DB.Model(EventLog{}).Where("chain_name = ?", chain).Last(&lastEvent); result.Error != nil {
		return nil, result.Error
	}
	return &lastEvent, nil
}
