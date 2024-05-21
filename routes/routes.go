package routes

import (
	"api-wa/app/controller"
	"api-wa/app/middleware"

	"github.com/gin-gonic/gin"
)

func InitRoute(app *gin.Engine, userController *controller.UserController) {
	route := app
	// Anda dapat menambahkan rute lainnya di sini sesuai kebutuhan aplikasi
	route.Use(middleware.JWTMiddleware())
	// Mendaftarkan rute untuk endpoint /register
	route.POST("/register", userController.RegisterUser)
	route.PUT("/update/:id", userController.UpdateUser)
	route.GET("/find/:id", userController.FindById)
	route.GET("/users", userController.FindAll)
	route.DELETE("/user/delete/:id", userController.DeleteUser)


}
