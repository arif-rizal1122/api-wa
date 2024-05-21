package routes

import (
	"api-wa/app/controller"
	"api-wa/app/middleware"

	"github.com/gin-gonic/gin"
)

func InitRoute(app *gin.Engine, userController *controller.UserController) {
	route := app

	// Mendaftarkan rute untuk endpoint /register
	// route.POST("/register", userController.RegisterUser)
	// route.PUT("/update/:id", userController.UpdateUser)
	// route.GET("/find/:id", userController.FindById)
	// route.GET("/users", userController.FindAll)
	// route.DELETE("/user/delete/:id", userController.DeleteUser)

	cms := route.Group("/cms", middleware.JWTMiddleware())
	{
	cms.POST("/register", userController.RegisterUser)
	cms.PUT("/update/:id", userController.UpdateUser)
	cms.GET("/find/:id", userController.FindById)
	cms.GET("/users", userController.FindAll)
	cms.DELETE("/user/delete/:id", userController.DeleteUser)
	}

	
}