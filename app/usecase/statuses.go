package usecase

import (
	"api-wa/app/domain/contract"
	"api-wa/app/domain/entity"
	"api-wa/app/domain/types/request"
	"api-wa/app/domain/types/response"
	"io"
	"net/http"
	"os"
	"path/filepath"
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

	filePath := filepath.Join("public", "img", filepath.Base(data.Picture))
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		inputFile, err := os.Open(data.Picture)
		if err != nil {
			return nil, err
		}
		defer inputFile.Close()

		outputFile, err := os.Create(filePath)
		if err != nil {
			return nil, err
		}
		defer outputFile.Close()

		_, err = io.Copy(outputFile, inputFile)
		if err != nil {
			return nil, err
		}
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
    status, err := s.repository.FindById(Idstatus)
    if err != nil {
        return err
    }

    // Hapus file gambar dari direktori
    if status.Picture != "" {
        filePath := filepath.Join("public", "img", filepath.Base(status.Picture))
        if _, err := os.Stat(filePath); err == nil {
            os.Remove(filePath)
        }
    }

    // Hapus status dari database
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
    if data.Picture != "" {
        if status.Picture != "" && status.Picture != data.Picture {
            oldFilePath := filepath.Join("public", "img", filepath.Base(status.Picture))
            if _, err := os.Stat(oldFilePath); err == nil {
                os.Remove(oldFilePath)
            }
        }
        filePath := filepath.Join("public", "img", filepath.Base(data.Picture))
        if _, err := os.Stat(filePath); os.IsNotExist(err) {
            inputFile, err := os.Open(data.Picture)
            if err != nil {
                return err
            }
            defer inputFile.Close()

            outputFile, err := os.Create(filePath)
            if err != nil {
                return err
            }
            defer outputFile.Close()

            _, err = io.Copy(outputFile, inputFile)
            if err != nil {
                return err
            }
        }
        status.Picture = data.Picture
    }

    errUpdate := s.repository.Update(status)
    if errUpdate != nil {
        return errUpdate
    }
    return nil
}
