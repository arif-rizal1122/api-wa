package service

import (
	"api-wa/app/domain/contract"
	"api-wa/app/domain/entity"
	"api-wa/app/helper"
	"api-wa/app/domain/types"
	"errors"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	Repository contract.UserRepository
}

func NewUserServiceImpl(repo contract.UserRepository) *UserService {
	return &UserService{Repository: repo}
}










// REGISTER USER FIXED
func (s *UserService) RegisterUser(data types.RequestUserRegister) (*helper.Payload, error) {
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

	response := helper.NewAPIregisterResponse(http.StatusOK, "User registered successfully", helper.ResponseUserRegister{
		Name:     createdUser.Name,
		Username: createdUser.Username,
		Email:    createdUser.Email,
		Password: string(hashedPassword),
		Phone:    createdUser.Phone,
	})

	return &response, nil
}




// UPDATE USER FIXED
func (s *UserService) UpdateUser(id int, data types.RequestUpdateUser)  error {

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




// FIND BY ID FIXED
func (s *UserService) FindById(Id int) (*helper.PayloadFind, error) {
    user, err := s.Repository.FindById(Id)
    if err != nil {
        return nil, err
    }
    if user == nil {
        return nil, errors.New("pengguna tidak ditemukan")
    }

    response := helper.ResponseFind{
        Name:     user.Name,
        Username: user.Username,
        Email:    user.Email,
        Phone:    user.Phone,
    }

	payload := &helper.PayloadFind{
		Message: "Find data success",
		Status: http.StatusOK,
		Data: response,
	}

    return payload, nil
}







// FIND ALL FIXED
func (s *UserService) FindAll() (*[]helper.PayloadFinds, error) {
    users, err := s.Repository.FindAll()
    if err != nil {
        return nil, err
    }

    var userResponses []helper.ResponseFinds
    for _, user := range *users {
        userResponses = append(userResponses, helper.ResponseFinds{
            Name:     user.Name,
            Username: user.Username,
            Email:    user.Email,
            Phone:    user.Phone,
        })
    }

    payload := &helper.PayloadFinds{
        Message: "Get all users success",
        Status:  http.StatusOK,
        Datas:   userResponses,
    }

    return &[]helper.PayloadFinds{*payload}, nil
}




// DELETED FIXED
func (s *UserService) DeleteUser(Id int) error {
	_, err := s.Repository.FindById(Id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("user not found") 
	} else if err != nil {
		return err 
	}

	err = s.Repository.DeleteUser(Id)
	if err != nil {
		return err
	}

	return nil
}







// func (s *UserService) LoginUser(data input.LoginUser) (input.ResponseUserLogin, error) {
// 	// Implementasi logika login pengguna di sini
// 	// Anda dapat menambahkan logika login sesuai kebutuhan aplikasi Anda
// 	return input.ResponseUserLogin{}, nil
// }


// func (s *UserService) GetAllUsers() ([]input.UserResponse, error) {
// 	// Implementasi logika pengambilan semua pengguna di sini
// 	// Anda dapat menambahkan logika pengambilan semua pengguna sesuai kebutuhan aplikasi Anda
// 	return []input.UserResponse{}, nil
// }
