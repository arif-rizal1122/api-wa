package middleware

import (
	"api-wa/app/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenStr := ctx.GetHeader("Authorization")

		if tokenStr == "" {
			ctx.JSON(http.StatusUnauthorized, helper.NewErrorsResponse("Unauthorized", http.StatusUnauthorized, "Token tidak ditemukan"))
			ctx.Abort()
			return
		}

		userId, role, userCurrent, err := helper.ValidateToken(tokenStr)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, helper.NewErrorsResponse("Unauthorized", http.StatusUnauthorized, err.Error()))
			ctx.Abort()
			return
		}


		// Menyimpan ID pengguna dalam konteks
		ctx.Set("userId", userId)
		ctx.Set("role", role)
		ctx.Set("userCurrent", userCurrent)
		// Melanjutkan eksekusi handler selanjutnya

		
		ctx.Next()
	}
}

