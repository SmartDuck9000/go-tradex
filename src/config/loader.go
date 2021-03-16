package config

import (
	"os"
	"strconv"
	"time"
)

type ConfigurationLoader interface {
	ReadConfig() *ServiceConfig
}

type EnvLoader struct{}

func (loader EnvLoader) ReadConfig() *ServiceConfig {
	return &ServiceConfig{
		DB: DBConfig{
			URL:             getEnv("DB_URL", ""),
			MaxIdleConn:     getIntEnv("MAX_IDLE_CONN", 10),
			MaxOpenConn:     getIntEnv("MAX_OPEN_CONN", 100),
			ConnMaxLifetime: getHoursEnv("CONN_MAX_LIFETIME", 1),
		},
		Host: getEnv("HOST", ""),
		Port: getEnv("PORT", ""),
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

// Simple helper function to read an environment variable into integer or return a default value
func getIntEnv(name string, defaultVal int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultVal
}

// Simple helper function to read an environment variable into time.Hour or return a default value
func getHoursEnv(name string, defaultVal int) time.Duration {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return time.Hour * time.Duration(value)
	}

	return time.Hour * time.Duration(defaultVal)
}
