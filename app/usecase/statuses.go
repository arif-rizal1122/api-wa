package usecase

import (
	"api-wa/app/domain/contract"
	"api-wa/app/domain/entity"
	"api-wa/app/domain/types/request"
	"api-wa/app/domain/types/response"
	"net/http"
)

type StatusUsecase struct {
	repository contract.StatusRepository
}

func NewStatusUsecase(repository contract.StatusRepository) *StatusUsecase {
	return &StatusUsecase{repository: repository}
}




func (s *StatusUsecase) CreateStatus(data request.RequestCreateStatus, userId int) (*response.PayloadStatusCreate, error) {
	status := &entity.Status{
		Caption: data.Caption,
		Picture: data.Picture,
	}
	result, err := s.repository.CreateStatus(status, userId)
	if err != nil {
		return nil, err
	}
	response := response.NewStatusCreateResponse("status created successfully", http.StatusCreated, response.StatusCreateResponse{
		Picture: result.Picture,
		Caption: result.Caption,
		UserId:  result.UserId,
	})
	return &response, nil
}



func (s *StatusUsecase) FindById(statusId int) (*response.PayloadStatusFind, error) {
	result, errFind := s.repository.FindById(statusId)
	if errFind != nil {
		return nil, errFind
	}
	response := response.NewStatusResponseFind("GET_DATA_SUCCESS", http.StatusOK, response.StatusResponseFind{
		Picture: result.Picture,
		Caption: result.Caption,
		UserId:  result.UserId,
	})
	return &response, nil
}




func (s *StatusUsecase) FindAll(userId int) (*response.PayloadStatusFinds, error) {
	results, errs := s.repository.FindAll(userId)
	if errs != nil {
		return nil, errs
	}
	var filteredResults []response.StatusResponseFinds
	for _, result := range *results {
		if result.UserId == userId {
			filteredResults = append(filteredResults, response.StatusResponseFinds{
				Picture: result.Picture,
				Caption: result.Caption,
				UserId:  result.UserId,
			})
		}
	}
	res := response.PayloadStatusFinds{
		Message:             "success get all",
		Status:              http.StatusOK,
		StatusResponseFinds: filteredResults,
	}
	return &res, nil
}




func (s *StatusUsecase) Delete(Idstatus int) error {
	errDelete := s.repository.Delete(Idstatus)
	if errDelete != nil {
		return errDelete
	}
	return nil
}


func (s *StatusUsecase) Update(id int, data request.RequestUpdateStatus) error {
	status, err := s.repository.FindById(id)
	if err != nil {
		return err
	}
	status.Caption = data.Caption
	status.Picture = data.Picture
	errUpdate := s.repository.Update(status)
	if errUpdate != nil {
		return errUpdate
	}
	return nil
}
