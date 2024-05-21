package contract

import (
	"api-wa/app/domain/entity"
	"api-wa/app/helper"
	"api-wa/app/domain/types"

	"github.com/gin-gonic/gin"
)

type UserRepository interface {
	Create(data *entity.User)         (*entity.User, error)
	Update(data *entity.User)         error
	FindById(Id int)                  (*entity.User, error)
	FindAll()                         (*[]entity.User, error)
	DeleteUser(Id int)                error
}


type UserService interface {
	RegisterUser(data types.RequestUserRegister)   (*entity.User, error)
	UpdateUser(Id int, data entity.User)           (string, error)
	FindById(Id int)                               (*helper.ResponseFind, error)
	FindAll()                                      (*[]helper.ResponseFinds, error)
    DeleteUser(Id int)							   (string, error)	
}

type UserController interface {
	RegisterUser(ctx *gin.Context)
	LoginUser(ctx *gin.Context)
	IsEmailAvailable(ctx *gin.Context)
	GetUserById(ctx *gin.Context)
	FindAll(ctx *gin.Context)
}
