package usecase

import (
	"api-wa/app/domain/contract"
	"api-wa/app/domain/types/request"
	"api-wa/app/domain/types/response"
	"api-wa/app/helper"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type AuthUsecaseUser struct {
	repository contract.UserRepository
}

func NewAuthUsecaseUser(repository contract.UserRepository) *AuthUsecaseUser {
	return &AuthUsecaseUser{
		repository: repository,
	}
}








func (s *UserUsecase) LoginUser(data request.AuthUserLoginRequest) (*response.ResponseUserLogin, error) {
	user, err := s.Repository.UserLogin(data.Email)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password)); err != nil {
		return nil, errors.New("wrong credentials")
	}

	token, err := helper.GenerateToken(user)
	if err != nil {
		return nil, err
	}

	response := &response.ResponseUserLogin{
		Email: user.Email,
		Token: token,
	}
	return response, nil
}
