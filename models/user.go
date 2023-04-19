package models

import (
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

func Login(username, password string) bool {
	for _, u := range UserList {
		if u.Username == username && u.Password == password {
			return true
		}
	}
	return false
}

func Register(username, password, email string) bool {
	if username == "" || password == "" || email == "" {
		return false
	}
	uu := User{generateUserUuid(), username, password, email}
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
