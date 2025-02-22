// Package: app
// Description: Contains application-level operations for managing and processing API interactions,
// including the integration with SWOClient and utility functions.
package app

import (
	"manager/models"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

// operations...
type operations struct {
	logger *zap.SugaredLogger
	db     *gorm.DB
}

// newOperations initializes operations with necessary configs
func newOperations(log *zap.SugaredLogger, db *gorm.DB) *operations {
	return &operations{
		logger: log,
		db:     db,
	}
}

func (o *operations) ProcessCreate(task models.Task) models.Response {
	o.logger.Info(task)
	return models.Response{
		StatusCode: 200,
		Message:    "Task created successfully",
	}

}
