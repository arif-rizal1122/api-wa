package routes

import (
	"api-wa/app/controller"
	"api-wa/app/middleware"

	"github.com/gin-gonic/gin"
)

func InitRoute(app *gin.Engine, userController *controller.UserController, statusController *controller.StatusController) {
	route := app



	route.POST("/register", userController.RegisterUser)
	route.POST("/login", userController.LoginUser)

	cms := route.Group("/cms", middleware.JWTMiddleware())
	{
		cms.GET("/users", userController.FindAll)
		cms.PUT("/update/:id", userController.UpdateUser)
		cms.GET("/find/:id", userController.FindById)
		cms.DELETE("/user/delete/:id", userController.DeleteUser)
		
		cms.POST("/status/create", statusController.CreateStatus)
	    cms.GET("/status/find/:id", statusController.FindById)
		cms.GET("/status/finds", statusController.FindAll)
		cms.POST("/status/delete/:id", statusController.Delete)
		cms.POST("/status/update/:id", statusController.Update)
	}


	
	
}