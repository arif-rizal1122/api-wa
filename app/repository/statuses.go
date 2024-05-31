package repository

import (
	"api-wa/app/domain/contract"
	"api-wa/app/domain/entity"
	"errors"

	"gorm.io/gorm"
)

type StatusRepository struct {
	DB *gorm.DB
}

func NewStatusRepository(db *gorm.DB) contract.StatusRepository {
	return &StatusRepository{DB: db}
}

func (r *StatusRepository) CreateStatus(data *entity.Status, userId int) (*entity.Status, error) {
	tx := r.DB.Begin()
	data.UserId = userId
	if err := tx.Table("statuses").Create(&data).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	return data, nil
}



func (r *StatusRepository) FindById(statusId int) (*entity.Status, error) {
	var status entity.Status
	tx := r.DB.Begin()
	if err := tx.First(&status, "id = ?", statusId).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	return &status, nil
}


func (r *StatusRepository) FindAll(userId int) (*[]entity.Status, error) {
	var statuses []entity.Status

	tx := r.DB.Begin()
	if err := tx.Where("user_id = ?", userId).Find(&statuses).Error; err != nil {
		tx.Rollback()
		return nil, errors.New("failed to retrieve statuses")
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	return &statuses, nil
}




func (r *StatusRepository) Update(data *entity.Status) error {
	tx := r.DB.Begin()
	if err := tx.Table("statuses").Save(&data).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}
	return nil
}


func (r *StatusRepository) Delete(statusId int) error {
	tx := r.DB.Begin()
	var status entity.Status
	if err := tx.First(&status, statusId).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Delete(&status).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

