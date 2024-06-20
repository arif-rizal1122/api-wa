package bootstrap

import (
	"api-wa/app/controller"
	"api-wa/app/repository"
	"api-wa/app/usecase"
	"api-wa/config"
	"api-wa/config/appconfig"
	"api-wa/database"
	"api-wa/routes"
	"log"
 _ "github.com/go-sql-driver/mysql"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func BootstrapApp() {

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

	// initialize route user
	userRepository := repository.NewUserRepository(database.DB)
	userusecase := usecase.NewUserUsecaseImpl(userRepository)
	userController := controller.NewUserController(*userusecase)
	// initialize auth
	authUsecaseUser := usecase.NewAuthUsecaseUser(userRepository)
	authController  := controller.NewAuthController(*authUsecaseUser)
	// initialize route status
	statusRepository := repository.NewStatusRepository(database.DB)
	statususecase    := usecase.NewStatusUsecase(statusRepository)
	statusController := controller.NewStatusController(*statususecase)

	// inject user controller to routes
	routes.InitRoute(app, userController, statusController, authController)

	// run app
	app.Run(":" + appconfig.PORT)
}
