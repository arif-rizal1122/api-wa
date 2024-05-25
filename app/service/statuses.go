package service

import (
	"api-wa/app/domain/contract"
	"api-wa/app/domain/entity"
	"api-wa/app/domain/types/request"
	"api-wa/app/domain/types/response"
	"net/http"
)

type StatusService struct {
	repository  	contract.StatusRepository
}

func NewStatusService(repository contract.StatusRepository) *StatusService {
	return &StatusService{repository: repository}
}




func (s *StatusService) Create(data   request.RequestCreateStatus) (*response.PayloadStatusCreate, error) {

	 var user entity.User
     status  := entity.Status{
		Caption: data.Caption,
		UserId: user.ID,
		Picture: data.Picture,
	 }

	 result, err := s.repository.Create(&status)
	 if err != nil {
		return nil, err
	 }

	 response := response.NewStatusCreateResponse("status created successfully", http.StatusOK, response.StatusCreateResponse{
		Picture: result.Picture,
		Caption: result.Caption,
		UserId: result.UserId,
	 })

	 return response, nil
}