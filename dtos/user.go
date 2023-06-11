package dtos

type UserLoginDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserRegisterDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type UserDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Roles    string `json:"roles"`
}
