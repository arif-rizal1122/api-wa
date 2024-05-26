package controller

import (
	"api-wa/app/domain/types/request"
	"api-wa/app/helper"
	"api-wa/app/service"

	"net/http"

	"github.com/gin-gonic/gin"
)

type StatusController struct {
	service service.StatusService
}

func NewStatusController(service service.StatusService) *StatusController {
	return &StatusController{service: service}
}



// Create handles the creation of a new status.
func (c *StatusController) CreateStatus(ctx *gin.Context) {
	var input request.RequestCreateStatus

	// Mengikat data JSON
	if err := ctx.ShouldBindJSON(&input); err != nil {
		response := helper.NewErrorsResponse("BAD_REQUEST", http.StatusUnprocessableEntity, err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response, err := c.service.CreateStatus(ctx, input)
	if err != nil {
		errRes := helper.NewErrorsResponse("BAD_REQUEST", http.StatusBadRequest, err.Error())
		ctx.JSON(http.StatusBadRequest, errRes)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": response.Message, "data": response.Data})
}






func (c *StatusController) Update(ctx *gin.Context) {
	
}

func (c *StatusController) FindById(ctx *gin.Context) {
	
}

func (c *StatusController) FindAll(ctx *gin.Context) {
	
}

func (c *StatusController) Delete(ctx *gin.Context) {
	
}
