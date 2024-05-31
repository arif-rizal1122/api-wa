package routes

import (
	"api-wa/app/controller"
	"api-wa/app/middleware"

	"github.com/gin-gonic/gin"
)

func InitRoute(app *gin.Engine, userController *controller.UserController, statusController *controller.StatusController, authController *controller.AuthController) {
	route := app

	route.POST("/register", userController.RegisterUser)
	route.POST("/login", userController.LoginUser)

	cms := route.Group("/cms", middleware.JWTMiddleware())
	{
		cms.GET("/user/finds", userController.FindAll)
		cms.DELETE("/user/delete/:id", userController.DeleteUser)
		cms.PUT("/user/update/:id", userController.UpdateUser)
		cms.GET("/user/find/:id", userController.FindById)
		
		cms.POST("/status/create", statusController.CreateStatus)
	    cms.GET("/status/find/:id", statusController.FindById)
		cms.GET("/status/finds", statusController.FindAll)
		cms.DELETE("/status/delete/:id", statusController.Delete)
		cms.PUT("/status/update/:id", statusController.Update)
	}
	
}