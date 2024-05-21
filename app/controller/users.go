package controller

import (
	"api-wa/app/domain/input"
	"api-wa/app/helper"
	"api-wa/app/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service service.UserService
}

func NewUserController(service service.UserService) *UserController {
	return &UserController{service: service}
}



func (c *UserController) RegisterUser(ctx *gin.Context) {
	var input input.RequestUserRegister
	if err := ctx.ShouldBindJSON(&input); err != nil {

		response := helper.NewErrorsResponse("Register user failed", http.StatusUnprocessableEntity, err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response, err := c.service.RegisterUser(input)
	if err != nil {
		errRes := helper.NewErrorsResponse("BAD_REQUEST", http.StatusBadRequest, err.Error())
		ctx.JSON(http.StatusBadRequest, errRes)
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": response.Message, "user": response.Data})
}




func (c *UserController) UpdateUser(ctx *gin.Context) {
	var input input.RequestUpdateUser

	if err := ctx.ShouldBindJSON(&input); err != nil {
		response := helper.NewErrorsResponse("Updated user failed", http.StatusUnprocessableEntity, err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response := helper.NewErrorsResponse("user id no found", http.StatusNotFound, err.Error())
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	err = c.service.UpdateUser(id, input)
	if err != nil {
		response := helper.NewErrorsResponse("Internal server error", http.StatusBadRequest, err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Updated success"})
}

func (c *UserController) FindById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response := helper.NewErrorsResponse("user id no found", http.StatusNotFound, err.Error())
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	user, err := c.service.FindById(id)
	if err != nil {
		response := helper.NewErrorsResponse("user id no found", http.StatusNotFound, err.Error())
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	ctx.JSON(http.StatusOK, user)
}








func (c *UserController) DeleteUser(ctx *gin.Context) {
	Id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response := helper.NewErrorsResponse("user id not found", http.StatusNotFound, err.Error())
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	err = c.service.DeleteUser(Id)
	if err != nil {
		response := helper.NewErrorsResponse("error deleting user", http.StatusBadRequest, err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "user deleted successfully"})
}

func (c *UserController) FindAll(ctx *gin.Context) {
	user, err := c.service.FindAll()
	if err != nil {
		response := helper.NewErrorsResponse("user id no found", http.StatusNotFound, err.Error())
		ctx.JSON(http.StatusNotFound, response)
		return
	}
	ctx.JSON(http.StatusOK, user)
}
