package models

import (
	"beego-car-rental/dtos"
	"errors"

	"github.com/google/uuid"
)

var (
	UserList map[string]*User
)

func init() {
	UserList = make(map[string]*User)

	uuid := generateUserUuid()
	u := User{uuid, "test", "1234", "test@gmail.com"}
	UserList[uuid] = &u
}

type User struct {
	Uuid     string `json:"uuid"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func AddUser(u User) string {
	if u.Email == "" || u.Password == "" || u.Username == "" {
		return ""
	}
	u.Uuid = generateUserUuid()
	UserList[u.Uuid] = &u
	return u.Uuid
}

func GetUser(uuid string) (u *User, err error) {
	if u, ok := UserList[uuid]; ok {
		return u, nil
	}
	return nil, errors.New("User does not exist")
}

func GetAllUsers() map[string]*User {
	return UserList
}

func UpdateUser(uuid string, uu *User) (a *User, err error) {
	if u, ok := UserList[uuid]; ok {
		if uu.Username != "" {
			u.Username = uu.Username
		}
		if uu.Password != "" {
			u.Password = uu.Password
		}
		return u, nil
	}
	return nil, errors.New("User does not exist")
}

func Login(userCreds dtos.UserLoginDto) bool {
	for _, u := range UserList {
		if u.Username == userCreds.Username && u.Password == userCreds.Password {
			return true
		}
	}
	return false
}

func Register(user dtos.UserRegisterDto) bool {
	if user.Username == "" || user.Password == "" || user.Email == "" {
		return false
	}
	uu := User{generateUserUuid(), user.Username, user.Password, user.Email}
	UserList[uu.Uuid] = &uu
	return true
}

func DeleteUser(uuid string) bool {
	userPresent := false
	if _, ok := UserList[uuid]; ok {
		userPresent = true
	}
	delete(UserList, uuid)
	return userPresent
}

func generateUserUuid() string {
	uuid := uuid.New()
	return uuid.String()
}
