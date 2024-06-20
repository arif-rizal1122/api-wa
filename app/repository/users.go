package repository

import (
	"api-wa/app/domain/contract"
	"api-wa/app/domain/entity"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) contract.UserRepository {
	return &UserRepository{DB: db}
}

func (u *UserRepository) Create(data *entity.User) (*entity.User, error) {
	tx := u.DB.Begin()
	if err := tx.Table("users").Create(&data).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	return data, nil
}

func (u *UserRepository) Update(data *entity.User) error {
	tx := u.DB.Begin()
	if err := tx.Table("users").Save(&data).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}
	return nil
}



func (u *UserRepository) DeleteUser(Id int) error {
	tx := u.DB.Begin()
	user, err := u.FindById(Id)
	if err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Delete(&user).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func (u *UserRepository) FindAll() (*[]entity.User, error) {
	var users []entity.User
	tx := u.DB.Begin()
	if err := tx.Find(&users).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	return &users, nil
}

func (u *UserRepository) FindById(Id int) (*entity.User, error) {
	var user entity.User
	tx := u.DB.Begin()
	result := tx.First(&user, "id = ?", Id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			tx.Rollback()
			return nil, errors.New("user not found")
		}
		tx.Rollback()
		return nil, fmt.Errorf("error finding user by ID: %w", result.Error)
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	return &user, nil
}

func (u *UserRepository) UserLogin(email string) (*entity.User, error) {
	var user entity.User
	tx := u.DB.Begin()
	result := tx.First(&user, "Email = ?", email)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			tx.Rollback()
			return nil, errors.New("credentials errors")
		}
		tx.Rollback()
		return nil, fmt.Errorf("error finding user by email: %w", result.Error)
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	return &user, nil
}



func (u *UserRepository) FindByUsername(username string) (*entity.User, error) {
	var usernameUser entity.User
    tx    :=  u.DB.Begin()
	result  :=  tx.First(&usernameUser, "username = ?", username)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			tx.Rollback()
			return nil, errors.New("credentials errors")
		}
		tx.Rollback()
		return nil, fmt.Errorf("error finding user by username: %w", result.Error)
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	return &usernameUser, nil
}