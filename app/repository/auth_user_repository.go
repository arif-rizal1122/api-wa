package repository

import (
	"api-wa/app/domain/entity"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

func (u *UserRepositoryctx) FindByEmail(email string) (*entity.User, error) {
	var user entity.User

	result := u.DB.First(&user, "Email = ?", email)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			// Email tidak ditemukan, return user kosong dan error user not found
			return &entity.User{}, errors.New("email not found")
		}
		// Handle other potential errors more gracefully (e.g., logging)
		return &entity.User{}, fmt.Errorf("error finding user by email: %w", result.Error)
	}

	// Email ditemukan, return user dan error email already exists
	return &user, errors.New("email already exists")
}
