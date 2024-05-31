package helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthUserCurrent(ctx *gin.Context) (string, error) {
	userCurrent, ok := ctx.Get("userCurrent")
	if !ok {
		errRes := NewErrorsResponse("UNAUTHORIZED", http.StatusUnauthorized, "access denied")
		ctx.JSON(http.StatusUnauthorized, errRes)
		return "", nil
	}
	userCurrentStr, ok := userCurrent.(string)
	if !ok {
		errRes := NewErrorsResponse("UNAUTHORIZED", http.StatusUnauthorized, "access denied")
		ctx.JSON(http.StatusUnauthorized, errRes)
		return "", nil
	}
	return userCurrentStr, nil
}


