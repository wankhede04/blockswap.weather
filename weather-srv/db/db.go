package db

import (
	"log"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gorm_logger "gorm.io/gorm/logger"
)

// StorageConfig contains configurations for storage, postgreSQL
type PostgresStorageConfig struct {
	URL        string // DataBase URL for connection
	DBDriver   string // DataBase driver
	DBHOST     string // DataBase host
	DBPORT     int64
	DBSSL      string // DataBase sslmode
	DBName     string // DataBase name
	DBUser     string // DataBase's user
	DBPassword string // User's password
}

type PostgresDataBase struct {
	DB     *gorm.DB
	Logger *logrus.Logger
}

func InitialMigration(dbURL string, logger logrus.Logger) (*PostgresDataBase, error) {
	gormConfig := &gorm.Config{Logger: gorm_logger.Default.LogMode(gorm_logger.Silent), DisableForeignKeyConstraintWhenMigrating: true}

	db, err := gorm.Open(postgres.Open(dbURL), gormConfig)
	if err != nil {
		log.Panicf("failed to connect database %s", err)
	}

	sqlDB, err := db.DB()

	// run migrations
	if err := db.AutoMigrate(&Membership{}, &WeatherReport{}, &EventLog{}); err != nil {
		log.Panicf("failed to automigrate tables %s", err)
	}

	sqlDB.SetMaxOpenConns(10) // Set the maximum number of open connections

	return &PostgresDataBase{DB: db, Logger: &logger}, err
}
