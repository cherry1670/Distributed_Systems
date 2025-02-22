package migration

import (
	"manager/models"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

func MigrateDBEntities(logger *zap.SugaredLogger, db *gorm.DB) error {
	// Run the migrations
	if err := db.AutoMigrate(&models.Task{}, &models.Task_Queue{}); err != nil {
		logger.Error("Failed to migrate database entities: ", err)
		return err
	}
	logger.Info("Database migration completed successfully.")
	return nil
}
