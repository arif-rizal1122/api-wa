package controller

import (
	"api-wa/app/domain/types/request"
	"api-wa/app/helper"
	"api-wa/app/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	usecase usecase.AuthUsecaseUser
}

func NewAuthController(usecase usecase.AuthUsecaseUser) *AuthController {
	return &AuthController{
		usecase: usecase,
	}
}






func (c *UserController) LoginUser(ctx *gin.Context) {
	var request request.AuthUserLoginRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.NewErrorsResponse("Bad Request", http.StatusBadRequest, err.Error()))
		return
	}

	response, err := c.usecase.LoginUser(request)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, helper.NewErrorsResponse("Unauthorized", http.StatusUnauthorized, err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, response)
}
