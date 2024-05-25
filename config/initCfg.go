package config


import (
	"api-wa/config/appconfig"
	"api-wa/config/dbconfig"
)





func InitConfig() {
	appconfig.InitAppConfig()
	dbconfig.InitDBConfig()
}