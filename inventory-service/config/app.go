package config

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
