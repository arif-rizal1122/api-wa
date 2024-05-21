package appconfig

import "os"

var PORT = ":8080"

func InitAppConfig() {

	if portEnv := os.Getenv("APP_PORT"); portEnv != "" {
		PORT = portEnv
	}
}
