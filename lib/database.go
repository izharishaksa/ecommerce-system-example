package lib

import (
	"time"
)

type DatabaseConfig struct {
	DSN                 string
	MaxIdleConnections  int
	MaxOpenConnections  int
	MaxIdleDuration     time.Duration
	MaxLifeTimeDuration time.Duration
}

func getDatabaseConfig() DatabaseConfig {
	return DatabaseConfig{
		DSN:                 getStringOrPanic("DB_DSN"),
		MaxIdleConnections:  getIntOrDefault("DB_MAX_IDLE_CONNECTIONS", 20),
		MaxOpenConnections:  getIntOrDefault("DB_MAX_OPEN_CONNECTIONS", 100),
		MaxIdleDuration:     time.Duration(getIntOrDefault("DB_MAX_IDLE_DURATION_IN_SECS", 60)) * time.Second,
		MaxLifeTimeDuration: time.Duration(getIntOrDefault("DB_MAX_LIFE_TIME_DURATION_IN_SECS", 300)) * time.Second,
	}
}
