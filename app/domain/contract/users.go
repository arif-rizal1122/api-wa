package contract

import (
	"api-wa/app/domain/entity"
	"api-wa/app/domain/types/response"
	"api-wa/app/domain/types/request"

	"github.com/gin-gonic/gin"
)


type UserRepository interface {
	Create(data *entity.User)         (*entity.User, error)
	Update(data *entity.User)         error
	FindById(Id int)                  (*entity.User, error)
	FindAll()                         (*[]entity.User, error)
	DeleteUser(Id int)                error
	FindByEmail(email string)         (*entity.User, error)
	UserLogin(email string)	          (*entity.User, error)
}


type UserService interface {
	RegisterUser(data request.RequestUserRegister)   (*response.ResponseUserRegister, error)
	UpdateUser(Id int, data entity.User)           (string, error)
	FindById(Id int)                               (*response.ResponseFind, error)
	FindAll()                                      (*[]response.ResponseFinds, error)
    DeleteUser(Id int)							   (string, error)
	
	LoginUser(data request.AuthUserLoginRequest) (*entity.User, error)
	FindByEmail(email string)                  (*entity.User, error)
}


type UserController interface {
	RegisterUser(ctx *gin.Context)
	UpdateUser(ctx *gin.Context)
	FindById(ctx *gin.Context)
	FindByEmail(ctx *gin.Context)
	DeleteUser(ctx *gin.Context)
	LoginUser(ctx *gin.Context)
}
