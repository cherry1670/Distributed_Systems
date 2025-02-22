// Package: api
// Description: This file contains the implementation of the API server, including its setup,
//              initialization, and route configurations. The server uses Gin for routing
//              and includes dependencies for logging and error handling.

package api

import (
	"context"
	"net/http"
	"sync"

	"manager/app"
	"manager/configs"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// APIServer Struct Encapsulates the components of the API server
type APIServer struct {
	addr      string
	ctx       context.Context
	wg        *sync.WaitGroup
	logger    *zap.SugaredLogger
	errorChan chan error
	Runtime   *http.Server
	DB        *gorm.DB
}

// NewAPIServer Function
// Initializes the APIServer struct with required dependencies
func NewAPIServer(ctx context.Context, wg *sync.WaitGroup, logger *zap.SugaredLogger,
	serverErr chan error, db *gorm.DB) *APIServer {
	// Initialize the server with the provided context, wait group, and logger
	return &APIServer{
		addr:      configs.GetConfig().Serverport,
		ctx:       ctx,
		wg:        wg,
		logger:    logger,
		errorChan: serverErr,
		DB:        db,
	}
}

// StartAPIServer Method
// Sets up the router, initializes routes, and starts the server
func (server *APIServer) StartAPIServer() {
	// Set Gin to release mode
	gin.SetMode(gin.ReleaseMode)

	// Initialize the handler with the required dependencies
	handler := app.NewHandler(server.ctx, server.wg, server.logger, server.DB)

	// Set up the main router and v1 API group
	routerEngine := gin.Default()

	// Register routes using the router setup function from business/router
	handler.SetupRoutes(routerEngine) // Pass handler to SetupRoutes

	// Configure the runtime HTTP server with the Gin router
	server.Runtime = &http.Server{Addr: server.addr, Handler: routerEngine}
	server.logger.Info("APP is now available at http://localhost", server.addr)

	// Start the server asynchronously
	server.wg.Add(1)
	go func() {
		defer server.wg.Done()
		if err := server.Runtime.ListenAndServe(); err != http.ErrServerClosed {
			server.errorChan <- err
		}
	}()
}
