package services

import (
	"beego-car-rental/dtos"
	"beego-car-rental/models"
	"errors"

	"github.com/google/uuid"
)

func LoginUser(userInfo dtos.UserLoginDto) bool {
	return true
}

func RegisterUser(registerInfo dtos.UserRegisterDto) error {
	if registerInfo.Username == "" || registerInfo.Password == "" || registerInfo.Email == "" {
		return errors.New("empty field, please fill it")
	}
	// todo: use builder pattern
	user := &models.User{Username: registerInfo.Username,
		Password: registerInfo.Password,
		Email:    registerInfo.Email,
		Uuid:     generateUserUuid(),
	}

	return user.Insert()
}

func LogoutUser() bool {
	return true
}

func generateUserUuid() string {
	uuid := uuid.New()
	return uuid.String()
}
