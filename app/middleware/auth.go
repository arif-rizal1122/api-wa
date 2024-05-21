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
			// Mengirim respons HTTP dengan status 401 (Unauthorized)
			ctx.JSON(http.StatusUnauthorized, helper.NewErrorsResponse("Unauthorized", http.StatusUnauthorized, "Token tidak ditemukan"))
			ctx.Abort()
			return
		}

		userId, err := helper.ValidateToken(tokenStr)
		if err != nil {
			// Mengirim respons HTTP dengan status 401 (Unauthorized) dan pesan error
			ctx.JSON(http.StatusUnauthorized, helper.NewErrorsResponse("Unauthorized", http.StatusUnauthorized, err.Error()))
			ctx.Abort()
			return
		}
		// Menyimpan ID pengguna dalam konteks
		ctx.Set("userId", userId)

		// Melanjutkan eksekusi handler selanjutnya
		ctx.Next()
	}
}
