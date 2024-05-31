package helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthAdminRole(ctx *gin.Context) (string, error) {
	role, ok := ctx.Get("role")
	if !ok {
		errRes := NewErrorsResponse("UNAUTHORIZED", http.StatusUnauthorized, "User role not found in context")
		ctx.JSON(http.StatusUnauthorized, errRes)
		return "", nil
	}
	roleStr, ok := role.(string)
	if !ok {
		errRes := NewErrorsResponse("UNAUTHORIZED", http.StatusUnauthorized, "User role not found in context")
		ctx.JSON(http.StatusUnauthorized, errRes)
		return "", nil
	}
	return roleStr, nil
}


