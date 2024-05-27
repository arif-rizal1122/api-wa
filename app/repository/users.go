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
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	err := tx.Table("users").Create(&data).Error
	if err != nil {
		tx.Rollback() 
		return nil, err
	}
	tx.Commit() 
	return data, nil
}

func (u *UserRepositoryctx) Update(data *entity.User) error {
	tx := u.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	err := tx.Table("users").Save(&data).Error
	if err != nil {
		tx.Rollback()
		return err
	}
 
	tx.Commit()
	return nil
}


func (u *UserRepositoryctx) DeleteUser(Id int) error {
	tx := u.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	user, err := u.FindById(Id)
	if err != nil {
		tx.Rollback() 
		return err
	}

	result := tx.Delete(&user)
	if result.Error != nil {
		tx.Rollback() 
		return result.Error
	}

	tx.Commit()
	return nil
}





func (u *UserRepositoryctx) FindAll() (*[]entity.User, error) {
	var users []entity.User

	tx := u.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	result := tx.Find(&users)
	if result.Error != nil {
		tx.Rollback()
		return nil, result.Error
	}

	tx.Commit()
	return &users, nil
}




func (u *UserRepositoryctx) FindById(Id int) (*entity.User, error) {
	var user entity.User

	tx := u.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	result := tx.First(&user, "id = ?", Id) 

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			tx.Rollback()
			return &entity.User{}, errors.New("user not found")
		}
		tx.Rollback()
		return &entity.User{}, fmt.Errorf("error finding user by ID: %w", result.Error)
	}

	tx.Commit()
	return &user, nil
}



func (u *UserRepositoryctx) UserLogin(email string) (*entity.User, error) {
	var user entity.User
	tx := u.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	result := tx.First(&user, "Email = ?", email)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			tx.Rollback()
			return nil, errors.New("credentials errors")
		}
		    tx.Rollback()
		return nil, fmt.Errorf("error finding user by email: %w", result.Error)
	}

	tx.Commit()
	return &user, nil
}