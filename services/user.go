package services

import (
	"beego-car-rental/dtos"
	"beego-car-rental/models"
	"beego-car-rental/utils"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// todo: better error messages
// todo: use builder pattern

func GetUser(user *models.User, id int) error {
	if id < 1 {
		return errors.New("id is less than one")
	}
	user.Id = id
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

func UpdateUser(userDto *dtos.UserDto, id int) error {
	if id < 1 {
		return errors.New("id is less than one")
	}
	user := &models.User{Id: id}
	if err := user.Read(); err != nil {
		return err
	}
	user.Email = userDto.Email
	user.Username = userDto.Username
	user.Password = userDto.Password
	if err := user.Update(); err != nil {
		return err
	}
	return nil
}

func DeleteUser(id int) error {
	user := &models.User{Id: id}
	if id < 1 {
		return errors.New("id is less than one")
	}
	if err := user.Delete(); err != nil {
		return err
	}
	return nil
}

func LoginUser(userInfo *dtos.UserLoginDto) error {
	if userInfo.Username == "" || userInfo.Password == "" {
		return errors.New("empty field, please fill it")
	}

	user := &models.User{Username: userInfo.Username}
	if err := user.Read("username"); err != nil {
		return err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInfo.Password)); err != nil {
		return err
	}
	// set session
	return nil
}

func RegisterUser(registerInfo *dtos.UserDto) error {
	if registerInfo.Username == "" || registerInfo.Password == "" || registerInfo.Email == "" {
		return errors.New("empty field, please fill it")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerInfo.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// todo: check regex for email and username
	user := &models.User{Username: registerInfo.Username,
		Password: string(hashedPassword),
		Email:    registerInfo.Email,
		Uuid:     utils.GenerateUuid(),
	}

	return user.Insert()
}

func LogoutUser() error {
	// remove session
	return nil
}
