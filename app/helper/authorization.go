package helper

import "github.com/gin-gonic/gin"



func AuthUserID(ctx *gin.Context) (int, error) {
	userId, ok := ctx.Get("userId")
	if !ok {
		return 0, nil
	}

	userIdInt, ok := userId.(int)
	if !ok {
		return 0, nil
	}
	return userIdInt, nil
}