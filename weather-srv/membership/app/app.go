package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	weatherService "github.com/wankhede04/blockswap.weather/weather-srv/weather-service"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// App ...
type App struct {
	logger         *logrus.Logger
	engine         *gin.Engine
	server         *http.Server
	weatherservice *weatherService.WeatherService
}

// NewApp is initializes the app
func NewApp(logger *logrus.Logger, addr string, weatherservice *weatherService.WeatherService) *App {
	r := gin.Default()
	srv := &http.Server{
		Addr:    addr,
		Handler: r,
	}
	return &App{logger: logger, engine: r, server: srv, weatherservice: weatherservice}
}

// Run the app on it's router
// Run the app on its router
func (a *App) Run() {
	// Create a wait group to wait for goroutines to finish
	var wg sync.WaitGroup
	wg.Add(3)

	// Start the report-weather handler in a goroutine
	go func() {
		defer wg.Done()
		a.engine.POST("/report-weather", a.weatherservice.AuthenticateMiddleware(), a.weatherservice.RateLimitMiddleware(), a.weatherservice.ReportWeatherHandler)
	}()

	// Start the server in a goroutine
	go func() {
		defer wg.Done()
		if err := a.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	// Start the weather service in a goroutine
	go func() {
		defer wg.Done()
		a.weatherservice.Run()
	}()

	a.logger.Infof("Weather Service has started. Press ctrl + C to exit.")

	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	// Create a context with a timeout for graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Shutdown the server
	if err := a.server.Shutdown(ctx); err != nil {
		a.logger.Errorf("Server shutdown error: %v", err)
	}

	// Close the weather service
	a.weatherservice.Close()

	// Wait for goroutines to finish
	wg.Wait()

	a.logger.Infoln("Weather Service has stopped")
}
