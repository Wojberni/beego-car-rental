package dtos

type UserLoginDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
