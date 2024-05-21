package bootstrap

import (
	"api-wa/app/controller"
	"api-wa/app/repository"
	"api-wa/app/service"
	"api-wa/config"
	"api-wa/config/appconfig"
	"api-wa/database"
	"api-wa/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func BootstrapApp() {
	// load env file
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	// init configuration (appConfig & dbConfig)
	config.InitConfig()

	// database connection
	database.ConnectDB()

	// init gin engine
	app := gin.Default()

	// init user repository, service, and controller
	userRepository := repository.NewUserRepository(database.DB)
	userService := service.NewUserServiceImpl(userRepository)

	// Mengubah tipe userService ke *service.UserServiceImpl
	userController := controller.NewUserController(*userService)

	// inject user controller to routes
	routes.InitRoute(app, userController)

	// run app
	app.Run(":" + appconfig.PORT)
}
