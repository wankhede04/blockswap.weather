package main

import (
	"github.com/wankhede04/blockswap.weather/weather-srv/config"
	weatherservice "github.com/wankhede04/blockswap.weather/weather-srv/weather-service"
	"github.com/wankhede04/blockswap.weather/weather-srv/worker"

	"github.com/wankhede04/blockswap.weather/weather-srv/membership/app"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sirupsen/logrus"
)

// toWorkerConfig converts the worker configuration from the application's config package to the worker.WorkerConfig.
func toWorkerConfig(config config.WorkerConfig) worker.WorkerConfig {
	return worker.WorkerConfig{
		ChainName:            config.ChainName,
		Provider:             config.Provider,
		RegistrationContract: config.RegistrationContract,
		StartBlockHeight:     config.StartBlockHeight,
	}
}

func main() {
	logger := logrus.New()

	// Read the application configuration
	cfg := config.NewViperConfig()

	// Read the database configuration from the application config
	postgresDbConfig := cfg.ReadDBConfig()
	dbURL := postgresDbConfig.AsPostgresDbUrl()

	// Read the service URL from the application config
	srvURL := cfg.ReadServiceConfig()

	// Read the worker configurations from the application config and convert to worker.WorkerConfig
	workersCfg := cfg.ReadWorkersConfig()
	workerConfigs := toWorkerConfig(workersCfg)

	// Create a new instance of the WeatherService
	weatherservice, err := weatherservice.NewWeatherService(dbURL, logger, &workerConfigs, 10)
	if err != nil {
		logger.Panicf("Unable to create weather service %s", err.Error())
	}
	defer weatherservice.Close()

	// Create a new instance of the application
	app := app.NewApp(logger, srvURL, weatherservice)

	// Run the application
	app.Run()
}
