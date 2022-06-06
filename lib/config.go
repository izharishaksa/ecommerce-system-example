package lib

import (
	"fmt"
	"github.com/spf13/viper"
	"time"
)

type Config struct {
	App      AppConfig
	Database DatabaseConfig
	Kafka    string
}

type AppConfig struct {
	Name     string
	HTTPPort int
}

func getAppConfig() AppConfig {
	return AppConfig{
		Name:     getStringOrPanic("APP_NAME"),
		HTTPPort: getIntOrPanic("APP_HTTP_PORT"),
	}
}

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

func LoadConfigByFile(path, fileName, fileType string) Config {
	viper.SetConfigName(fileName)
	viper.SetConfigType(fileType)
	viper.AddConfigPath(path)
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Error reading config file, use automatic environment instead %s\n", err)
	}

	return Config{
		App:      getAppConfig(),
		Database: getDatabaseConfig(),
		Kafka:    getStringOrPanic("KAFKA_BROKERS"),
	}
}
