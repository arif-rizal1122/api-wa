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
	return &StatusController{
		service: service,
	}
}




func (c *StatusController) Create(ctx *gin.Context) {
	var input request.RequestCreateStatus
	if err := ctx.ShouldBindJSON(&input); err != nil {
		response := helper.NewErrorsResponse("add status by id success", http.StatusUnprocessableEntity, err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response, err := c.service.Create(input)
	if err != nil {
		errRes := helper.NewErrorsResponse("BAD_REQUEST", http.StatusBadRequest, err.Error())
		ctx.JSON(http.StatusBadRequest, errRes)
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": response.Message, "data": response.Data})
}