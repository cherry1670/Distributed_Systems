package configs

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type configs struct {
	// Database configuration
	Host       string `env:"DB_HOST" envDefault:"localhost"`
	Port       int64  `env:"DB_PORT" envDefault:"5432"`
	Username   string `env:"DB_USERNAME" envDefault:"postgres"`
	Password   string `env:"DB_PASSWORD" envDefault:"password"`
	DbName     string `env:"DB_NAME" envDefault:"taskmanager"`
	Serverport string `env:"SERVER_PORT" envDefault:"8080"`
	LogLevel   string `env:"LOG_LEVEL" envDefault:"info"`
}

var config *configs

func (c *configs) InitConfigs(envFile string) (*configs, error) {
	// if envfile is mentioned, and doesn't point to a valid location.
	if envFile != "" {
		err := godotenv.Load(envFile)
		if err != nil {
			return nil, fmt.Errorf("failed to load env file from configured path %v", err)
		}
	}
	config = &configs{
		Host:       getEnv("DB_HOST", "localhost"),
		Port:       getEnvAsInt("DB_PORT", 5432),
		Username:   getEnv("DB_USERNAME", "postgres"),
		Password:   getEnv("DB_PASSWORD", "password"),
		DbName:     getEnv("DB_NAME", "taskmanager"),
		Serverport: getEnv("SERVER_PORT", "8080"),
		LogLevel:   getEnv("LOG_LEVEL", "info"),
	}
	return config, nil
}

// GetConfig returns the initialized config instance.
func GetConfig() *configs {
	return config
}

// getEnv gets the env by key or return the default.
func getEnv(key, defaultVal string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultVal
}

// getEnvAsInt gets the env by key or use the default and convert into int64.
func getEnvAsInt(key string, defaultVal int64) int64 {
	if strValue, ok := os.LookupEnv(key); ok {
		intValue, err := strconv.ParseInt(strValue, 10, 64)
		if err != nil {
			return defaultVal
		}
		return intValue
	}
	return defaultVal
}
