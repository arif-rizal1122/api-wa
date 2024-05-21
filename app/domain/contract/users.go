package contract

import (
	"api-wa/app/domain/entity"
	"api-wa/app/helper"
	"api-wa/app/domain/input"

	"github.com/gin-gonic/gin"
)

type UserRepository interface {
	Create(data *entity.User)         (*entity.User, error)
	Update(data *entity.User)         error
	FindById(Id int)                  (*entity.User, error)
	FindAll()                         (*[]entity.User, error)
	DeleteUser(Id int)                error

	
	FindByEmail(email string)		  (*entity.User, error)
	FindByUsername(username string)   (*entity.User, error)
}


type UserService interface {
	RegisterUser(data input.RequestUserRegister)   (*entity.User, error)
	UpdateUser(Id int, data entity.User)           (string, error)
	FindById(Id int)                               (*helper.ResponseFind, error)
	FindAll()                                      (*[]input.UserResponse, error)
    DeleteUser(Id int)							   (string, error)


	
	LoginUser(data input.LoginUser)                (input.ResponseUserLogin, error)
	IsEmailAvailable(data input.CheckEmailUser)    error
}

type UserController interface {
	RegisterUser(ctx *gin.Context)
	LoginUser(ctx *gin.Context)
	IsEmailAvailable(ctx *gin.Context)
	GetUserById(ctx *gin.Context)
	FindAll(ctx *gin.Context)
}
