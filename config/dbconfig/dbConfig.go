package dbconfig

import "os"

var (
	DB_DRIVER   = "mysql"
	DB_HOST     = "127.0.0.1"
	DB_PORT     = "3306"
	DB_NAME     = "WA-API"
	DB_USER     = "root"
	DB_PASSWORD = ""
)

func InitDBConfig() {
	if dbDriver := os.Getenv("DB_DRIVER"); dbDriver != "" {
		DB_DRIVER = dbDriver
	}
	if dbHost := os.Getenv("DB_HOST"); dbHost != "" {
		DB_HOST = dbHost
	}
	if dbPort := os.Getenv("DB_PORT"); dbPort != "" {
		DB_PORT = dbPort
	}
	if dbName := os.Getenv("DB_NAME"); dbName != "" {
		DB_NAME = dbName
	}
	if dbUser := os.Getenv("DB_USER"); dbUser != "" {
		DB_USER = dbUser
	}
	if dbPassword := os.Getenv("DB_PASSWORD"); dbPassword != "" {
		DB_PASSWORD = dbPassword
	}
	
}
