package services

import (
	"beego-car-rental/dtos"
	"beego-car-rental/models"
	"beego-car-rental/utils"
	"errors"
	"net/mail"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

func LoginUser(userInfo *dtos.UserLoginDto) (string, error) {
	if userInfo.Username == "" || userInfo.Password == "" {
		return "", errors.New("empty field, please fill it")
	}

	user := &models.User{Username: userInfo.Username}
	if err := user.Read("username"); err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInfo.Password)); err != nil {
		return "", err
	}
	return user.Uuid, nil
}

func RegisterUser(registerInfo *dtos.UserDto) error {
	if err := validateRegisterData(registerInfo); err != nil {
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerInfo.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &models.User{
		Username: registerInfo.Username,
		Password: string(hashedPassword),
		Email:    registerInfo.Email,
		Uuid:     utils.GenerateUuid(),
	}

	return user.Insert()
}

func LogoutUser() error {
	return nil
}

func validateRegisterData(registerInfo *dtos.UserDto) error {
	if registerInfo.Username == "" || registerInfo.Password == "" || registerInfo.Email == "" {
		return errors.New("empty field, please fill it")
	}

	if _, err := mail.ParseAddress(registerInfo.Email); err != nil {
		return err
	}

	validUsername, _ := regexp.MatchString("[a-zA-Z0-9]{3,}", registerInfo.Username)
	if !validUsername {
		return errors.New("invalid username, 3 or more alphanumeric signs required")
	}

	validPassword, _ := regexp.MatchString("[a-zA-Z]{10,}", registerInfo.Password)
	if !validPassword {
		return errors.New("invalid password, 10 or more alphanumeric signs required")
	}

	return nil
}
