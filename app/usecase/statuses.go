package usecase

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

type StatusUsecase struct {
	repository contract.StatusRepository
}

func NewStatusUsecase(repository contract.StatusRepository) *StatusUsecase {
	return &StatusUsecase{repository: repository}
}


func (s *StatusUsecase) CreateStatus(ctx *gin.Context, data request.RequestCreateStatus) (*response.PayloadStatusCreate, error) {
	userId, err := helper.ValidateUserIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	status := &entity.Status{
		Caption: data.Caption,
		Picture: data.Picture,
		UserId:  userId,
	}
	result, err := s.repository.CreateStatus(status)
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



func (s *StatusUsecase) FindById(ctx *gin.Context, statusId int) (*response.PayloadStatusFind, error) {
	
	userIdInt, err  :=  helper.ValidateUserIDFromContext(ctx)
	if err != nil {
		return nil, err
	}
	result, errFind     := s.repository.FindById(statusId)
	if errFind != nil {
		return nil, errFind
	}
	if userIdInt != result.UserId {
		return nil, errors.New("status id not found")
	}

	response   :=  response.NewStatusResponseFind("GET_DATA_SUCCESS", http.StatusOK, response.StatusResponseFind{
		Picture: result.Picture,
		Caption: result.Caption,
	})
	return &response, nil
}




func (s *StatusUsecase) FindAll(ctx *gin.Context) (*response.PayloadStatusFinds, error) {
	userId, err := helper.ValidateUserIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	results, errs := s.repository.FindAll()
	if errs != nil {
		return nil, errs
	}

	var filteredResults []response.StatusResponseFinds
	for _, result := range *results {
		if result.UserId == userId {
			filteredResults = append(filteredResults, response.StatusResponseFinds{
				Picture: result.Picture,
				Caption: result.Caption,
			})
		}
	}

	res := response.PayloadStatusFinds{
		Message: "success get all",
		Status: http.StatusOK,
		StatusResponseFinds: filteredResults,
	}
	return &res, nil
}




func (s *StatusUsecase) Delete(ctx *gin.Context, Idstatus int) error {
	userId, err   :=   helper.ValidateUserIDFromContext(ctx)
	if err != nil {
		return err
	}

	errDelete := s.repository.Delete(Idstatus)
	if errDelete != nil {
		return errDelete
	}
	if userId != Idstatus {
		return errors.New("status not found")
	}
	return nil
}





func (s *StatusUsecase) Update(ctx *gin.Context, data *request.RequestUpdateStatus) error {
	userId, err := helper.ValidateUserIDFromContext(ctx)
	if err != nil {
		return err
	}

	status, err := s.repository.FindById(data.StatusId)
	if err != nil {
		return err
	}

	if userId != status.UserId {
		return errors.New("status not found")
	}

	status.Caption = data.Caption
	status.Picture = data.Picture

	errUpdate := s.repository.Update(status)
	if errUpdate != nil {
		return errUpdate
	}
	return nil
}

