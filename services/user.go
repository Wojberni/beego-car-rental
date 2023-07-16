package services

import (
	"beego-car-rental/dtos"
	"beego-car-rental/models"
	"beego-car-rental/utils"
	"errors"
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

	hashedPassword, err := utils.GeneratePassword(userDto.Password)
	if err != nil {
		return err
	}

	user.Email = userDto.Email
	user.Username = userDto.Username
	user.Password = hashedPassword
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

func CreateUser(userDto *dtos.UserDto) error {
	hashedPassword, err := utils.GeneratePassword(userDto.Password)
	if err != nil {
		return err
	}

	user := &models.User{
		Uuid:     utils.GenerateUuid(),
		Username: userDto.Username,
		Password: string(hashedPassword),
		Email:    userDto.Email,
		// todo: Roles: ,
	}

	return user.Insert()
}
