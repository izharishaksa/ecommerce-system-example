package lib

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	App      AppConfig
	Database DatabaseConfig
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
	}
}
