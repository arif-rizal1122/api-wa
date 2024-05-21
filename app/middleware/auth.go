package middleware

import (
	"api-wa/app/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

// JWTMiddleware adalah middleware untuk otorisasi menggunakan JWT token
func JWTMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenStr := ctx.GetHeader("Authorization")
		if tokenStr == "" {
			helper.NewErrorsResponse("Unauthorized", http.StatusUnauthorized, "Token tidak ditemukan")
			ctx.Abort()
			return
		}

		userId, err := helper.ValidateToken(tokenStr)
		if err != nil {
			helper.NewErrorsResponse("Unauthorized", http.StatusUnauthorized, err.Error())
			ctx.Abort()
			return
		}

		ctx.Set("userId", userId)
		ctx.Next()
	}
}
