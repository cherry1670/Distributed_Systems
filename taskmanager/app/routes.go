package app

import (
	"context"
	"manager/middleware"
	"manager/models"
	"sync"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// Handler encapsulates dependencies for API requests.
type Handler struct {
	logger     *zap.SugaredLogger
	wg         *sync.WaitGroup // Shared WaitGroup for managing active connections
	ctx        context.Context
	operations models.Operations
}

// NewHandler initializes a new Handler.
func NewHandler(ctx context.Context, wg *sync.WaitGroup, log *zap.SugaredLogger, db *gorm.DB) *Handler {
	return &Handler{
		logger:     log,
		wg:         wg,
		ctx:        ctx,
		operations: newOperations(log, db),
	}
}

// SetupRoutes registers routes under the provided API group.
func (handler *Handler) SetupRoutes(router *gin.Engine) {
	advancedRoutes := router.Group("api/v1/")
	advancedRoutes.Use(middleware.RemoveJSONSuffixFromParams())
	{
		// advancedRoutes.GET("/task", handler.GetTasks)
		// advancedRoutes.GET("/task/:id", handler.GetTaskByID)
		advancedRoutes.POST("/task", handler.CreateTask)

	}

}
