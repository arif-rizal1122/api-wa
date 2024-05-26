package repository

import (
	"api-wa/app/domain/contract"
	"api-wa/app/domain/entity"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type UserRepositoryctx struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) contract.UserRepository {
	return &UserRepositoryctx{DB: db}
}

func (u *UserRepositoryctx) Create(data *entity.User) (*entity.User, error) {
	tx := u.DB.Begin() 
	err := tx.Table("users").Create(&data).Error
	if err != nil {
		tx.Rollback() // Rollback jika terjadi kesalahan
		return nil, err
	}

	tx.Commit() 
	return data, nil
}

func (u *UserRepositoryctx) Update(data *entity.User) error {
	tx := u.DB.Begin() 

	err := tx.Table("users").Save(&data).Error
	if err != nil {
		tx.Rollback() // Rollback jika terjadi kesalahan
		return err
	}

	tx.Commit()
	return nil
}


func (u *UserRepositoryctx) DeleteUser(Id int) error {
	tx := u.DB.Begin() 
	user, err := u.FindById(Id)
	if err != nil {
		tx.Rollback() // Rollback jika terjadi kesalahan
		return err
	}

	result := tx.Delete(&user)
	if result.Error != nil {
		tx.Rollback() // Rollback jika terjadi kesalahan
		return result.Error
	}

	tx.Commit()
	return nil
}






func (u *UserRepositoryctx) FindAll() (*[]entity.User, error) {
	var users []entity.User

	result := u.DB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return &users, nil
}




func (u *UserRepositoryctx) FindById(Id int) (*entity.User, error) {
	var user entity.User
	result := u.DB.First(&user, "id = ?", Id) // Use named placeholders for security

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return &entity.User{}, errors.New("user not found")
		}
		// Handle other potential errors more gracefully (e.g., logging)
		return &entity.User{}, fmt.Errorf("error finding user by ID: %w", result.Error)
	}
	return &user, nil
}



func (u *UserRepositoryctx) UserLogin(email string) (*entity.User, error) {
	var user entity.User

	result := u.DB.First(&user, "Email = ?", email)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			// Email tidak ditemukan, return nil dan tanpa error
			return nil, errors.New("credentials errors")
		}
		// Handle other potential errors more gracefully (e.g., logging)
		return nil, fmt.Errorf("error finding user by email: %w", result.Error)
	}

	// Email ditemukan, return data pengguna tanpa error
	return &user, nil
}