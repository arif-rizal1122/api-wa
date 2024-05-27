package helper

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)


func ValidateUserIDFromContext(ctx *gin.Context) (int, error) {
	userId, ok := ctx.Get("userId")
	if !ok {
		errRes := NewErrorsResponse("UNAUTHORIZED", http.StatusUnauthorized, "User ID not found in context")
		ctx.JSON(http.StatusUnauthorized, errRes)
		return 0, errors.New("user ID not found in context")
	}

	userIdInt, ok := userId.(int)
	if !ok {
		errRes := NewErrorsResponse("UNAUTHORIZED", http.StatusUnauthorized, "User ID is not of type int")
		ctx.JSON(http.StatusUnauthorized, errRes)
		return 0, errors.New("user ID is not of type int")
	}
	return userIdInt, nil
}
