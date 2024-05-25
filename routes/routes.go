package routes

import (
	"api-wa/app/controller"
	"api-wa/app/middleware"

	"github.com/gin-gonic/gin"
)

func InitRoute(app *gin.Engine, userController *controller.UserController) {
	route := app



	route.POST("/register", userController.RegisterUser)
	route.POST("/login", userController.LoginUser)


	cms := route.Group("/cms", middleware.JWTMiddleware())
	{
	cms.PUT("/update/:id", userController.UpdateUser)
	cms.GET("/find/:id", userController.FindById)
	cms.GET("/users", userController.FindAll)
	cms.DELETE("/user/delete/:id", userController.DeleteUser)
	}

	
}