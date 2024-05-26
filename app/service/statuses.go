package service

import (
	"api-wa/app/domain/contract"
	"api-wa/app/domain/entity"
	"api-wa/app/domain/types/request"
	"api-wa/app/domain/types/response"
	"api-wa/app/helper"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type StatusService struct {
	repository contract.StatusRepository
}

func NewStatusService(repository contract.StatusRepository) *StatusService {
	return &StatusService{repository: repository}
}


func (s *StatusService) CreateStatus(ctx *gin.Context, data request.RequestCreateStatus) (*response.PayloadStatusCreate, error) {
	// Get user ID from the JWT token (assuming you have a middleware that sets the user ID in the context)
	userId, ok := ctx.Get("userId")
	if !ok {
		errRes := helper.NewErrorsResponse("UNAUTHORIZED", http.StatusUnauthorized, "User ID not found in context")
		ctx.JSON(http.StatusUnauthorized, errRes)
		return nil, errors.New("user ID not found in context")
	}
	userIdInt, ok := userId.(int)
	if !ok {
		errRes := helper.NewErrorsResponse("UNAUTHORIZED", http.StatusUnauthorized, "User ID is not of type int")
		ctx.JSON(http.StatusUnauthorized, errRes)
		return nil, errors.New("user ID is not of type int")
	}

	

	status := &entity.Status{
		Caption: data.Caption,
		Picture: data.Picture,
		UserId:  userIdInt,
	}

	result, err := s.repository.CreateStatus(status)
	if err != nil {
		return nil, err
	}

	response := response.NewStatusCreateResponse("status created successfully", http.StatusOK, response.StatusCreateResponse{
		Picture: result.Picture,
		Caption: result.Caption,
		UserId:  result.UserId,
	})

	return &response, nil
}












func (s *StatusService) Update() {

}

func (s *StatusService) FindById() {

}

func (s *StatusService) FindAll() {

}

func (s *StatusService) Delete() {

}
