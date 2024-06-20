package controller

import (
	"api-wa/app/domain/types/request"
	"api-wa/app/helper"
	"api-wa/app/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	usecase usecase.UserUsecase
}

func NewUserController(usecase usecase.UserUsecase) *UserController {
	return &UserController{usecase: usecase}
}

/*
*fix
*
 */
func (c *UserController) RegisterUser(ctx *gin.Context) {
	var input request.RequestUserRegister

	if err := ctx.ShouldBindJSON(&input); err != nil {
		response := helper.NewErrorsResponse("ERROR", http.StatusUnprocessableEntity, "REGISTER USER FAILED")
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	response, err := c.usecase.RegisterUser(input)
	if err != nil {
		errRes := helper.NewErrorsResponse("ERROR", http.StatusBadRequest, "BAD_REQUEST")
		ctx.JSON(http.StatusBadRequest, errRes)
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": response.Message, "user": response.Data})
}

/*
*fix
*
 */

func (c *UserController) UpdateUser(ctx *gin.Context) {
	var input request.RequestUpdateUser
	if err := ctx.ShouldBindJSON(&input); err != nil {
		response := helper.NewErrorsResponse("ERROR", http.StatusUnprocessableEntity, "INVALID INPUT DATA")
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response := helper.NewErrorsResponse("ERROR", http.StatusBadRequest, "INVALID PARAMETER")
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	userCurrent, errAuthCurrent := helper.AuthUserCurrent(ctx)
	if errAuthCurrent != nil {
		response := helper.NewErrorsResponse("ERROR", http.StatusUnauthorized, "USER UN AUTHENTICATED")
		ctx.JSON(http.StatusUnauthorized, response)
		return
	}
	email, errFInd := c.usecase.FindByUsername(userCurrent)
	if errFInd != nil {
		response := helper.NewErrorsResponse("ERROR", http.StatusNotFound, "User email not found")
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	if email.ID != id {
		response := helper.NewErrorsResponse("ERROR", http.StatusForbidden, "User ID does not match authenticated user")
		ctx.JSON(http.StatusForbidden, response)
		return
	}
	err = c.usecase.UpdateUser(id, input)
	if err != nil {
		response := helper.NewErrorsResponse("ERROR", http.StatusInternalServerError, "Failed to update user, possible duplicate data")
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

/*
*fix
*
 */
func (c *UserController) FindById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response := helper.NewErrorsResponse("ERROR", http.StatusNotFound, "NOT FOUND")
		ctx.JSON(http.StatusNotFound, response)
		return
	}
	user, err := c.usecase.FindById(id)
	if err != nil {
		response := helper.NewErrorsResponse("ERROR", http.StatusBadRequest, "BAD REQUEST")
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	ctx.JSON(http.StatusOK, user)
}

/*
*fix
*
 */

func (c *UserController) DeleteUser(ctx *gin.Context) {
	role, _ := helper.AuthAdminRole(ctx)
	if role != "admin@gmail.com" { // Pastikan role yang diverifikasi adalah admin
		response := helper.NewErrorsResponse("ERROR", http.StatusForbidden, "FORBIDDEN ACCESS")
		ctx.JSON(http.StatusForbidden, response)
		return
	}
	Id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response := helper.NewErrorsResponse("ERROR", http.StatusNotFound, "NOT FOUND")
		ctx.JSON(http.StatusNotFound, response)
		return
	}
	err = c.usecase.DeleteUser(Id)
	if err != nil {
		response := helper.NewErrorsResponse("ERROR", http.StatusBadRequest, "BAD_REQUEST")
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "user deleted successfully"})
}

/*
*fix
*
 */

func (c *UserController) FindAll(ctx *gin.Context) {
	role, _ := helper.AuthAdminRole(ctx)
	if role != "admin@gmail.com" {
		response := helper.NewErrorsResponse("ERROR", http.StatusForbidden, "FORBIDDEN ACCESS")
		ctx.JSON(http.StatusForbidden, response)
		return
	}
	user, err := c.usecase.FindAll()
	if err != nil {
		response := helper.NewErrorsResponse("ERROR", http.StatusBadRequest, "FIND ALL ERROR")
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	ctx.JSON(http.StatusOK, user)
}
