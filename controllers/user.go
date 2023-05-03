package controllers

import (
	"beego-car-rental/dtos"
	"beego-car-rental/models"
	"beego-car-rental/services"
	"encoding/json"
	"fmt"

	beego "github.com/beego/beego/v2/server/web"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

// @Title CreateUser
// @Description Create user
// @Param 	body 	body 	models.User 	true 	"Body for user content"
// @Success 200 {string} Created user with Uuid!
// @Failure 403 {string} Data missing, please fill all data!
// @Accept json
// @router / [post]
func (u *UserController) Post() {
	user := new(models.User)
	var message string
	json.Unmarshal(u.Ctx.Input.RequestBody, user)
	if err := user.Insert(); err != nil {
		message = fmt.Sprintf("Error creating user: %v", err.Error())
	} else {
		message = fmt.Sprintf("Created user: %v!", user.Username)
	}

	u.Data["json"] = map[string]string{"message": message}
	u.ServeJSON()
}

// @Title GetAll
// @Description Get all Users
// @Success 200 {object} models.User
// @Failure 404 {string} Error retrieving data, please try again later!
// @Accept json
// @router / [get]
func (u *UserController) GetAll() {
	var users models.UserList
	users.ReadAll()
	u.Data["json"] = users
	u.ServeJSON()
}

// @Title Get
// @Description Get user by uuid
// @Param 	uuid 	path 	string 	true 	"The uuid of user to get"
// @Success 200 {object} models.User
// @Failure 403 {string} Uuid is empty!
// @Accept json
// @router /:uuid [get]
func (u *UserController) Get() {
	uid := u.GetString(":uuid")
	user := &models.User{Uuid: uid}
	if uid != "" {
		if err := user.Read("Uuid"); err != nil {
			u.Data["json"] = map[string]string{"message": err.Error()}
		} else {
			u.Data["json"] = user
		}
	}
	u.ServeJSON()
}

// @Title Update
// @Description Update the user
// @Param 	uuid 	path 	string 			true 	"The uuid you want to update"
// @Param	body 	body 	models.User 	true 	"Body for user content"
// @Success 200 {object} models.User
// @Failure 403 {string} Uuid is not int!
// @Accept json
// @router /:uuid [put]
func (u *UserController) Put() {
	uid := u.GetString(":uuid")
	if uid != "" {
		user := new(models.User)
		json.Unmarshal(u.Ctx.Input.RequestBody, user)
		if err := user.Update(); err != nil {
			u.Data["json"] = map[string]string{"message": err.Error()}
		} else {
			u.Data["json"] = user
		}
	}
	u.ServeJSON()
}

// @Title Delete
// @Description delete the user
// @Param	uuid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} Deleted user: uuid
// @Failure 403 {string} User not found: uuid
// @Accept json
// @router /:uuid [delete]
func (u *UserController) Delete() {
	uid := u.GetString(":uuid")
	user := &models.User{Uuid: uid}
	var message string
	if uid != "" {
		if err := user.Delete(); err != nil {
			message = fmt.Sprintf("User not found error: %s", err.Error())
		} else {
			message = fmt.Sprintf("Deleted user: %s", uid)
		}
	}
	u.Data["json"] = map[string]string{"message": message}
	u.ServeJSON()
}

// @Title Login
// @Description Logs user into the system
// @Param 	body 	body 	dtos.UserLoginDto 	true 	"Body of user login info"
// @Success 200 {string} Login success!
// @Failure 403 {string} User does not exist!
// @Accept json
// @router /login [post]
func (u *UserController) Login() {
	var userLogin dtos.UserLoginDto
	var message string
	json.Unmarshal(u.Ctx.Input.RequestBody, &userLogin)
	if services.LoginUser(userLogin) {
		message = fmt.Sprintf("Login success for user %v!", userLogin.Username)
	} else {
		message = "User does not exist!"
	}
	u.Data["json"] = map[string]string{"message": message}
	u.ServeJSON()
}

// @Title Register
// @Description Register user into the system
// @Param 	body 	body 	dtos.UserRegisterDto 	true 	"Body of user register info"
// @Success 200 {string} Register success!
// @Failure 403 {string} Register failure! Fill all fields!
// @Accept json
// @router /register [post]
func (u *UserController) Register() {
	var user dtos.UserRegisterDto
	var message string
	json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	if err := services.RegisterUser(user); err != nil {
		message = fmt.Sprintf("Register success for user %v!", user.Username)
	} else {
		message = fmt.Sprintf("Register failure! Error: %v", err.Error())
	}
	u.Data["json"] = map[string]string{"message": message}
	u.ServeJSON()
}

// @Title Logout
// @Description Logs out current logged in user session
// @Success 200 {string} Logout success!
// @Accept json
// @router /logout [post]
func (u *UserController) Logout() {
	u.Data["json"] = map[string]string{"message": "Logout success!"}
	u.ServeJSON()
}
