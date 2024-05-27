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


func NewStatusRepository(db *gorm.DB) contract.StatusRepository {
	return &StatusRepositoryctx{DB: db}
}


func (s *StatusRepositoryctx) CreateStatus(data *entity.Status) (*entity.Status, error) {
	tx := s.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	err := tx.Table("statuses").Create(&data).Error
	if err != nil {
		tx.Rollback() 
		return nil, err
	}
	tx.Commit() 
	return data, nil
}


func (s *StatusRepositoryctx) FindById(statusId int) (*entity.Status, error) {
	var status entity.Status
	tx := s.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	err := tx.First(&status, "id = ?", statusId).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return &status, nil
}



func (u *StatusRepositoryctx) FindAll() (*[]entity.Status, error) {
	var statuses []entity.Status
	tx := u.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	results := tx.Find(&statuses)
	if results.Error != nil {
		tx.Rollback()
		return nil, errors.New("failed to retrieve statuses")
	}
	tx.Commit()
	return &statuses, nil
}


func (s *StatusRepositoryctx) Update(data *entity.Status)  error {
	 tx := s.DB.Begin()
	 defer func() {
	 if r := recover(); r != nil {
		tx.Rollback()
	   }
    }()
	  err   	:= tx.Table("statuses").Save(&data).Error
	  if err != nil {
		   tx.Rollback()
		   return  err
	  }
	  tx.Commit()
	  return nil
}



func (s *StatusRepositoryctx) Delete(statusId int) error {
	tx := s.DB.Begin()
	defer func() {
	if r := recover(); r != nil {
	   tx.Rollback()
	   }
   }()
   status, err := s.FindById(statusId)
   if err != nil {
	   tx.Rollback()
	   return err
   }
	errDelete := tx.Delete(&status)
	if errDelete != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}


