package usecase

import (
	"api-wa/app/domain/contract"
	"api-wa/app/domain/entity"
	"api-wa/app/domain/types/request"
	"api-wa/app/domain/types/response"
	"errors"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserUsecase struct {
	Repository contract.UserRepository
}

func NewUserUsecaseImpl(repo contract.UserRepository) *UserUsecase {
	return &UserUsecase{Repository: repo}
}





func (s *UserUsecase) RegisterUser(data request.RequestUserRegister) (*response.Payload, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user := &entity.User{
		Name:      data.Name,
		Username:  data.Username,
		Email:     data.Email,
		Password:  string(hashedPassword),
		Phone:     data.Phone,
		CreatedAt: time.Now(),
	}
	createdUser, err := s.Repository.Create(user)
	if err != nil {
		return nil, err
	}
	response := response.NewAPIregisterResponse(http.StatusOK, "User registered successfully", response.ResponseUserRegister{
		Name:     createdUser.Name,
		Username: createdUser.Username,
		Email:    createdUser.Email,
		Phone:    createdUser.Phone,
	})
	return &response, nil
}





func (s *UserUsecase) UpdateUser(id int, data request.RequestUpdateUser) error {
	user, err := s.Repository.FindById(id)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("user not found")
	}
	if data.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		user.Password = string(hashedPassword)
	}
	user.Name = data.Name
	user.Username = data.Username
	user.Email = data.Email
	user.Phone = data.Phone
	user.UpdatedAt = time.Now()
	err = s.Repository.Update(user)
	if err != nil {
		return err
	}
	return nil
}




func (s *UserUsecase) FindById(Id int) (*response.PayloadFind, error) {
	user, err := s.Repository.FindById(Id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("pengguna tidak ditemukan")
	}
	responseFind := response.ResponseFind{
		Name:     user.Name,
		Username: user.Username,
		Email:    user.Email,
		Phone:    user.Phone,
	}
	payload := &response.PayloadFind{
		Message: "Data ditemukan dengan sukses",
		Status:  http.StatusOK,
		Data:    responseFind,
	}
	return payload, nil
}




func (s *UserUsecase) FindAll() (*response.PayloadFinds, error) {
	users, err := s.Repository.FindAll()
	if err != nil {
		return nil, err
	}
	var userResponses []response.ResponseFinds
	for _, user := range *users {
		userResponses = append(userResponses, response.ResponseFinds{
			Name:     user.Name,
			Username: user.Username,
			Email:    user.Email,
			Phone:    user.Phone,
		})
	}
	payload := &response.PayloadFinds{
		Message: "Get all users success",
		Status:  http.StatusOK,
		Datas:   userResponses,
	}
	return payload, nil
}




func (s *UserUsecase) DeleteUser(Id int) error {
	_, err := s.Repository.FindById(Id)
 	if err != nil {
	   return err
	}
	err = s.Repository.DeleteUser(Id)
	if err != nil {
		return err
	}
	return nil
}



func (s *UserUsecase) FindByEmail(email string) (*entity.User, error) {
    emailUser, err := s.Repository.FindByEmail(email)
    if err != nil {
        return nil, err
    }
    return emailUser, nil
}

