package services

import (
	"beego-car-rental/dtos"
	"beego-car-rental/models"
	"errors"

	"github.com/google/uuid"
)

func GetUser(user *models.User, uuid string) error {
	if uuid == "" {
		return errors.New("uuid is empty")
	}
	user.Uuid = uuid
	if err := user.Read(); err != nil {
		return err
	}
	return nil
}

func GetAllUsers(users *models.UserList) error {
	if err := users.ReadAll(); err != nil {
		return err
	}
	return nil
}

func UpdateUser(user *models.User, uuid string) error {
	if uuid == "" {
		return errors.New("uuid is empty")
	}
	user.Uuid = uuid
	if err := user.Read(); err != nil {
		return err
	}
	return nil
}

func DeleteUser(uuid string) error {
	user := &models.User{Uuid: uuid}
	if uuid == "" {
		return errors.New("uuid is empty")
	}
	if err := user.Delete(); err != nil {
		return err
	}
	return nil
}

func LoginUser(userInfo dtos.UserLoginDto) error {
	return nil
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

func LogoutUser() error {
	return nil
}

func generateUserUuid() string {
	uuid := uuid.New()
	return uuid.String()
}
