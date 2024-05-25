package repository

import (
	"api-wa/app/domain/contract"
	"api-wa/app/domain/entity"
	"errors"

	"gorm.io/gorm"
)

type StatusRepositoryctx struct {
	DB          *gorm.DB
}


func NewStatusRepositoryctx(db *gorm.DB) contract.StatusRepository {
	return &StatusRepositoryctx{DB: db}
}



func (s *StatusRepositoryctx) Create(data *entity.Status) (*entity.Status, error) {
	var user entity.User

    result := s.DB.Preload("User").Where("id = ?", data.UserId).First(&user)
    if err := result.Error; err != nil {
            return nil, errors.New("user not found")
    }

    if user.ID != 0 {
        data.UserId = user.ID 
        if err := s.DB.Create(&data).Error; err != nil {
            return nil, err
        }
        return data, nil
    }
    return nil, errors.New("user not found") 
}



func (s *StatusRepositoryctx) Update(data entity.Status)  error {
	err  := s.DB.Save(&data).Error
	if err != nil {
		return err
	}
	return nil
}


func (s *StatusRepositoryctx) FindById(statusId int) (*entity.Status, error) {
	var status entity.Status
	err := s.DB.First(&status, "id = ?", statusId).Error
	if err != nil {
		return nil, err
	}
	return &status, nil
}

func (u *StatusRepositoryctx) FindAll() (*[]entity.Status, error) {
	var statuses []entity.Status

	result := u.DB.Find(&statuses)
	if result.Error != nil {
		return nil, result.Error
	}
	return &statuses, nil
}



func (s *StatusRepositoryctx) Delete(statusId int) error {
	status, err := s.FindById(statusId)
	if err != nil {
		return err
	}

	errDelete := s.DB.Delete(&status)
	if errDelete != nil {
		return err
	}
	return nil
}


