package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	App      AppConfig
	Database DatabaseConfig
}

func Load() Config {
	viper.SetConfigName("application")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Error reading config file, use automatic environment instead %s\n", err)
	}

	return Config{
		App:      getAppConfig(),
		Database: getDatabaseConfig(),
	}
}
