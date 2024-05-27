package controller

import (
	"api-wa/app/domain/types/request"
	"api-wa/app/helper"
	"api-wa/app/usecase"
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"
)

type StatusController struct {
	Usecase usecase.StatusUsecase
}
func NewStatusController(Usecase usecase.StatusUsecase) *StatusController {
	return &StatusController{Usecase: Usecase}
}


func (c *StatusController) CreateStatus(ctx *gin.Context) {
	var input request.RequestCreateStatus

	if err := ctx.ShouldBindJSON(&input); err != nil {
		response := helper.NewErrorsResponse("BAD_REQUEST", http.StatusUnprocessableEntity, err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	response, err := c.Usecase.CreateStatus(ctx, input)
	if err != nil {
		errRes := helper.NewErrorsResponse("BAD_REQUEST", http.StatusBadRequest, err.Error())
		ctx.JSON(http.StatusBadRequest, errRes)
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": response.Message, "data": response.Data})
}





func (c *StatusController) FindById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response := helper.NewErrorsResponse("status id no found", http.StatusNotFound, err.Error())
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	status, err := c.Usecase.FindById(ctx, id)
	if err != nil {
		errRes := helper.NewErrorsResponse("BAD_REQUEST", http.StatusBadRequest, err.Error())
		ctx.JSON(http.StatusBadRequest, errRes)
		return
	}
	ctx.JSON(http.StatusOK, status)
}





func (c *StatusController) FindAll(ctx *gin.Context) {
	status, err := c.Usecase.FindAll(ctx)
	if err != nil {
		errRes := helper.NewErrorsResponse("BAD_REQUEST", http.StatusBadRequest, err.Error())
		ctx.JSON(http.StatusBadRequest, errRes)
		return
	}
	ctx.JSON(http.StatusOK, status)
}



func (c *StatusController) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response := helper.NewErrorsResponse("status id no found", http.StatusNotFound, err.Error())
		ctx.JSON(http.StatusNotFound, response)
		return
	}		

	errDelete := c.Usecase.Delete(ctx, id)
	if errDelete != nil {
		errRes := helper.NewErrorsResponse("BAD_REQUEST", http.StatusBadRequest, errDelete.Error())
		ctx.JSON(http.StatusBadRequest, errRes)
		return
	}
    ctx.JSON(http.StatusOK, gin.H{"message": "delete statuse by id success"})
}




func (c *StatusController) Update(ctx *gin.Context) {
	var input request.RequestUpdateStatus
	if err := ctx.ShouldBindJSON(&input); err != nil {
		response := helper.NewErrorsResponse("Updated status failed", http.StatusUnprocessableEntity, err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	_, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response := helper.NewErrorsResponse("user id no found", http.StatusNotFound, err.Error())
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	err = c.Usecase.Update(ctx, &input)
	if err != nil {
		response := helper.NewErrorsResponse("Internal server error", http.StatusBadRequest, err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Updated success success"})
}




