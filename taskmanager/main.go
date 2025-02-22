package main

import (
	"context"
	"flag"
	"log"
	"manager/cmd/api"
	"manager/configs"
	"manager/database"
	"manager/logger"
	"manager/migration"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	envFile := flag.String(
		"env-file",
		"",
		"runtime env config file path",
	)
	flag.Parse()

	// Runtime config.
	config, err := configs.GetConfig().InitConfigs(*envFile)
	if err != nil {
		log.Fatal(err)
	}
	// Custom logger used by application. It doesn't apply to GORM Library logs.
	logger := logger.NewLogger(config.LogLevel)

	db, err := database.InitDB(config.Host, config.Port, config.Username, config.DbName, config.Password)
	if err != nil {
		logger.Fatal(err)
	}
	err = migration.MigrateDBEntities(logger, db)
	if err != nil {
		logger.Fatal(err)
	}
	os.Exit(0)
	// Define runtime context and waitgroup
	runtimeContext, runtimeContextCancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	// Set up signal handling for graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	// Initializing  the Client instance, which is responsible for establishing the connections

	// Channel for capturing API server errors
	serverError := make(chan error, 1)
	// Initialize and start the API server
	apiServer := api.NewAPIServer(runtimeContext, &wg, logger, serverError, db)
	apiServer.StartAPIServer()

	// Handle application shutdown
	select {
	case err := <-serverError: // If the API server encounters an error
		logger.Errorf("Failed to run API server: %+v", err)
		runtimeContextCancel()
	case <-stop: // If an OS signal is received
		logger.Info("Received an OS signal, shutting down gracefully...")
		runtimeContextCancel()

		// Gracefully shutdown the API server
		if err := apiServer.Runtime.Shutdown(runtimeContext); err != nil {
			logger.Errorf("Failed to shutdown API server gracefully: %+v", err)
		}
	}
	wg.Wait()

}
