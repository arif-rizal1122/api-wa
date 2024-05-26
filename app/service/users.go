package service

import (
	"api-wa/app/domain/contract"
	"api-wa/app/domain/entity"
	"api-wa/app/domain/types/request"
	"api-wa/app/domain/types/response"
	"api-wa/app/helper"
	"errors"
	"fmt"
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











func (s *UserService) RegisterUser(data request.RequestUserRegister) (*response.Payload, error) {

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










func (s *UserService) UpdateUser(id int, data request.RequestUpdateUser)  error {

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

	fmt.Println(data.Password)

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




func (s *UserService) FindById(Id int) (*response.PayloadFind, error) {
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




func (s *UserService) FindAll() (*[]response.PayloadFinds, error) {
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

    return &[]response.PayloadFinds{*payload}, nil
}


func (s *UserService) DeleteUser(Id int) error {
	// Langkah 1: Periksa apakah pengguna ada
	_, err := s.Repository.FindById(Id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("user not found") // Pengguna tidak ditemukan
	} else if err != nil {
		return err // Error lain
	}

	// Langkah 2: Hapus pengguna berdasarkan ID
	err = s.Repository.DeleteUser(Id)
	if err != nil {
		return err
	}

	return nil // Pengguna berhasil dihapus
}



func (s *UserService) LoginUser(data request.AuthUserLoginRequest) (*response.ResponseUserLogin, error) {
	// Check email
	user, err := s.Repository.UserLogin(data.Email)
	if err != nil {
		return nil, err
	}

	// Check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password)); err != nil {
		return nil, errors.New("wrong credentials")
	}

	// Generate JWT token
	token, err := helper.GenerateToken(user)
	if err != nil {
		return nil, err
	}

	// Create response
	response := &response.ResponseUserLogin{
		Email: user.Email,
		Token: token,
	}

	return response, nil
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
