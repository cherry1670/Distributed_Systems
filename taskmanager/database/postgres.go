package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// GetDBLogger sets GORM logger properties.
func GetDBLogger() logger.Interface {
	return logger.New(

		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			// Slow SQL threshold
			SlowThreshold: 2 * time.Second,
			LogLevel:      logger.Silent,
			Colorful:      false,
		},
	)

}

// InitDB initializes postgres driver with configured performance parameters.
func InitDB(host string, port int64, user string, database string, password string) (*gorm.DB, error) {

	connStr := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s application_name=konnect"+
		" sslmode=disable connect_timeout=10",
		host, port, user, database, password)
	db, err := gorm.Open(
		postgres.New(
			postgres.Config{
				DSN:                  connStr,
				PreferSimpleProtocol: true,
			}),
		&gorm.Config{
			Logger: GetDBLogger(),
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %v", err)
	}

	return db, nil
}
