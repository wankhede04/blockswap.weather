package db

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Membership represents the membership model
type Membership struct {
	gorm.Model        // GORM model for common fields (ID, CreatedAt, UpdatedAt, DeletedAt)
	Address    string `gorm:"unique_index"` // Address of the membership (unique index)
	Status     string // Status of the membership
	LastCall   int64  // Last call timestamp for the membership
}

// MembershipStatus represents the possible status values for the membership
type MembershipStatus string

const (
	Unregistered MembershipStatus = "Unregistered"
	Registered   MembershipStatus = "Registered"
	Resigned     MembershipStatus = "Resigned"
)

// WeatherReport represents the weather report model
type WeatherReport struct {
	gorm.Model          // GORM model for common fields (ID, CreatedAt, UpdatedAt, DeletedAt)
	MembershipID uint   // ID of the associated membership
	Report       string // Weather report data
}

// EventLog represents the event log model
type EventLog struct {
	gorm.Model                // GORM model for common fields (ID, CreatedAt, UpdatedAt, DeletedAt)
	ID              int       // Custom ID for the event log
	CreatedAt       int64     // Creation timestamp of the log
	UpdatedAt       int64     // Last update timestamp of the log
	BlockHeight     uint64    // Block height of the event
	ChainName       string    // Name of the blockchain
	TransactionHash string    // Hash of the transaction
	Address         string    // Address associated with the event
	Timestamp       time.Time // Timestamp of the log
}
