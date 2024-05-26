package helper

import (
	"api-wa/app/domain/entity"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var SECRET_KEY = []byte("secretkey")

type JWTClaims struct {
	ID int `json:"id"`
	jwt.RegisteredClaims
}

// GenerateToken generates a new JWT token for the given user
func GenerateToken(user *entity.User) (string, error) {
	claims := JWTClaims{
		ID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(60 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(SECRET_KEY)
	return ss, err
}

// ValidateToken validates a JWT token and returns the user ID if valid
func ValidateToken(tokenStr string) (int, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &JWTClaims{}, func(t *jwt.Token) (interface{}, error) {
		return SECRET_KEY, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return 0, errors.New("invalid token signature")
		}
		return 0, errors.New("your token was expired")
	}

	claims, ok := token.Claims.(*JWTClaims)
	if !ok || !token.Valid {
		return 0, errors.New("your token was expired")
	}

	return claims.ID, nil
}
