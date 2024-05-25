package contract

import (
	"api-wa/app/domain/entity"
	"api-wa/app/domain/types/request"
	"api-wa/app/domain/types/response"

	"github.com/gin-gonic/gin"
)







type StatusRepository interface {
	Create(data *entity.Status)			 (*entity.Status, error)
	Update(data entity.Status) 			 error
	FindById(statusId int)               (*entity.Status, error)       
	FindAll()							 (*[]entity.Status, error)
	Delete(statusId  int)                error
}






type StatusService interface {
	Create(data         request.RequestCreateStatus)  (*response.PayloadStatusCreate, error)
	Update(data         request.RequestUpdateStatus)  (*response.PayloadUpdateStatus, error)	
	FindById(statusId int)							(*response.ResponseFind, error)
	FindAll()                                       (*[]response.ResponseFinds, error)
	Delete(statusId int)                            (string, error)
}





type StatusController interface {
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	FindById(ctx *gin.Context)
	FindAll(ctx *gin.Context)
	Delete(ctx *gin.Context)
}