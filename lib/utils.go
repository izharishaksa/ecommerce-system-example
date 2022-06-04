package lib

import (
	"github.com/spf13/viper"
	"os"
	"strconv"
)

func checkKey(key string) {
	_, envExists := os.LookupEnv(key)

	if !viper.IsSet(key) && !envExists {
		panic(key + " key is not set")
	}
}

func getString(key string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return viper.GetString(key)
}

func getStringOrPanic(key string) string {
	checkKey(key)
	return getString(key)
}

func getIntOrDefault(key string, def int) int {
	value := getString(key)

	if intValue, err := strconv.Atoi(value); err == nil {
		return intValue
	}

	return def
}

func getIntOrPanic(key string) int {
	checkKey(key)
	return getIntOrDefault(key, 0)
}
