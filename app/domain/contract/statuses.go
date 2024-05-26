package contract

import (
	"api-wa/app/domain/entity"
	"api-wa/app/domain/types/request"
	"api-wa/app/domain/types/response"

	"github.com/gin-gonic/gin"
)





type StatusRepository interface {
	CreateStatus(data *entity.Status)       (*entity.Status, error)
	Update(data *entity.Status)      		 error
	FindById(statusId int)          	    (*entity.Status, error)
	FindAll()                        	    (*[]entity.Status, error)
	Delete(statusId int)         			 error
}


type StatusService interface {
	Create(data         request.RequestCreateStatus)  (*response.StatusCreateResponse, error)
	Update(data         request.RequestUpdateStatus)  (*response.StatusUpdateResponse, error)	
	FindById(statusId int)							(*response.StatusResponseFind, error)
	FindAll()                                       (*[]response.StatusResponseFinds, error)
	Delete(statusId int)                            (string, error)
}





type StatusController interface {
	create(ctx *gin.Context)
	Update(ctx *gin.Context)
	FindById(ctx *gin.Context)
	FindAll(ctx *gin.Context)
	Delete(ctx *gin.Context)
}
