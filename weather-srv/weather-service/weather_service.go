package weatherservice

import (
	"sync"

	"github.com/wankhede04/blockswap.weather/weather-srv/db"
	"github.com/wankhede04/blockswap.weather/weather-srv/watcher"
	"github.com/wankhede04/blockswap.weather/weather-srv/worker"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type WeatherService struct {
	worker    *worker.Worker
	watcher   *watcher.WatcherSRV
	Database  *db.PostgresDataBase
	logger    *logrus.Logger
	dbPool    *db.ConnectionPool // Custom connection pool
	dbMutex   sync.RWMutex       // Mutex for database connection synchronization
	semaphore *sync.WaitGroup
}

func NewWeatherService(dbURL string, logger *logrus.Logger, cfg *worker.WorkerConfig, maxConcurrentConnections int) (*WeatherService, error) {
	database, err := db.InitialMigration(dbURL, *logger)
	if err != nil {
		return nil, err
	}

	wkr := worker.NewWorker(logger, *cfg, database)
	watcher, err := watcher.NewWatcherSRV(database, logger, wkr)
	if err != nil {
		return nil, err
	}

	// Create a connection pool with the specified maximum number of concurrent connections
	dbPool, err := db.NewConnectionPool(database.DB, maxConcurrentConnections)
	if err != nil {
		return nil, err
	}

	// Create a semaphore with the specified maximum number of concurrent connections
	semaphore := &sync.WaitGroup{}
	semaphore.Add(maxConcurrentConnections)

	return &WeatherService{
		worker:    wkr,
		watcher:   watcher,
		Database:  database,
		logger:    logger,
		dbPool:    dbPool,
		dbMutex:   sync.RWMutex{},
		semaphore: semaphore,
	}, nil
}

func (r *WeatherService) Run() {
	go r.watcher.Run()
}

func (r *WeatherService) getDBConnection() (*gorm.DB, error) {
	r.dbMutex.Lock()
	defer r.dbMutex.Unlock()

	db, err := r.dbPool.AcquireConnection()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (r *WeatherService) releaseDBConnection(db *gorm.DB) {
	r.dbMutex.Lock()
	defer r.dbMutex.Unlock()

	r.dbPool.ReleaseConnection(db)
}

func (r *WeatherService) Close() {
	r.dbMutex.Lock()
	defer r.dbMutex.Unlock()

	// Close all database connections in the connection pool
	r.dbPool.CloseConnections()
}
