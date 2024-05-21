package contract

import (
	"api-wa/app/domain/entity"
	"api-wa/app/domain/types"

	"github.com/gin-gonic/gin"
)



type AuthUserLoginRepository interface {
	LoginUser(data *entity.User)         (*entity.User, error)
	FindByEmail(email string)	    	 (*entity.User, error)
}


type AuthUserLoginService interface {
	LoginUser(data types.AuthUserLoginRequest) (*entity.User, error)
	FindByEmail(email string)                  (*entity.User, error)
}


type AuthUserLoginController interface {
	LoginUser(ctx *gin.Context)
}