package controllers

import (
	"beego-car-rental/dtos"
	"beego-car-rental/services"
	"encoding/json"
	"fmt"

	beego "github.com/beego/beego/v2/server/web"
)

type AuthController struct {
	beego.Controller
}

// @Title Login
// @Description Logs user into the system
// @Param 	body 	body 	dtos.UserLoginDto 	true 	"Body of user login info"
// @Success 200 {string} message: "Login success for user: username"
// @Failure 500 {string} error: "message"
// @Accept json
// @router /login [post]
func (u *AuthController) Login() {
	userLogin := &dtos.UserLoginDto{}
	json.Unmarshal(u.Ctx.Input.RequestBody, userLogin)
	if uuid, err := services.LoginUser(userLogin); err != nil {
		u.Data["json"] = map[string]string{"error": err.Error()}
		u.Ctx.Output.SetStatus(500)
	} else {
		params := []interface{}{uuid, userLogin.Username}
		u.SetSession("login", params)

		message := fmt.Sprintf("Login success for user %v!", userLogin.Username)
		u.Data["json"] = map[string]string{"message": message}
	}
	u.ServeJSON()
}

// @Title Register
// @Description Register user into the system
// @Param 	body 	body 	dtos.UserDto 	true 	"Body of user register info"
// @Success 200 {string} message: "Register success for user: username"
// @Failure 500 {string} error: "message"
// @Accept json
// @router /register [post]
func (u *AuthController) Register() {
	user := &dtos.UserRegisterDto{}
	json.Unmarshal(u.Ctx.Input.RequestBody, user)
	if err := services.RegisterUser(user); err != nil {
		u.Data["json"] = map[string]string{"error": err.Error()}
		u.Ctx.Output.SetStatus(500)
	} else {
		message := fmt.Sprintf("Register success for user %v!", user.Username)
		u.Data["json"] = map[string]string{"message": message}
	}

	u.ServeJSON()
}

// @Title Logout
// @Description Logs out current logged in user session
// @Success 200 {string} Logout success!
// @Accept json
// @router /logout [post]
func (u *AuthController) Logout() {
	if err := u.DelSession("login"); err != nil {
		u.Data["json"] = map[string]string{"error": err.Error()}
		u.Ctx.Output.SetStatus(500)
	} else {
		u.Data["json"] = map[string]string{"message": "Logout success!"}
	}
	u.ServeJSON()
}
