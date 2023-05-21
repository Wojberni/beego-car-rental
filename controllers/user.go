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

// @Title GetAll
// @Description Get all Users
// @Success 200 {object} models.User
// @Failure 404 {string} error: "message"
// @Accept json
// @router / [get]
func (u *UserController) GetAll() {
	users := &models.UserList{}
	if err := services.GetAllUsers(users); err != nil {
		u.Data["json"] = map[string]string{"error": err.Error()}
	} else {
		u.Data["json"] = users
	}
	u.ServeJSON()
}

// @Title Get
// @Description Get user by id
// @Param 	id 	path 	int 	true 	"The id of user to get"
// @Success 200 {object} models.User
// @Failure 403 {string} error: "message"
// @Accept json
// @router /:id [get]
func (u *UserController) Get() {
	id, _ := u.GetInt(":id")
	user := &models.User{}
	if err := services.GetUser(user, id); err != nil {
		u.Data["json"] = map[string]string{"error": err.Error()}
	} else {
		u.Data["json"] = user
	}
	u.ServeJSON()
}

// @Title Update
// @Description Update the user
// @Param 	id 	path 	int 			true 	"The id you want to update"
// @Param	body 	body 	dtos.UserDto 	true 	"Body for user content"
// @Success 200 {string} message: "Updated user: id"
// @Failure 403 {string} error: "message"
// @Accept json
// @router /:id [put]
func (u *UserController) Put() {
	id, _ := u.GetInt(":id")
	user := &dtos.UserDto{}
	json.Unmarshal(u.Ctx.Input.RequestBody, user)
	if err := services.UpdateUser(user, id); err != nil {
		u.Data["json"] = map[string]string{"error": err.Error()}
	} else {
		message := fmt.Sprintf("Updated user: %v", id)
		u.Data["json"] = map[string]string{"message": message}
	}
	u.ServeJSON()
}

// @Title Delete
// @Description delete the user
// @Param	id		path 	int	true		"The id you want to delete"
// @Success 200 {string} message: "Deleted user: id"
// @Failure 403 {string} error: "message"
// @Accept json
// @router /:id [delete]
func (u *UserController) Delete() {
	id, _ := u.GetInt(":id")
	if err := services.DeleteUser(id); err != nil {
		u.Data["json"] = map[string]string{"error": err.Error()}
	} else {
		message := fmt.Sprintf("Deleted user: %v", id)
		u.Data["json"] = map[string]string{"message": message}
	}
	u.ServeJSON()
}

// @Title Login
// @Description Logs user into the system
// @Param 	body 	body 	dtos.UserLoginDto 	true 	"Body of user login info"
// @Success 200 {string} message: "Login success for user: username"
// @Failure 403 {string} error: "message"
// @Accept json
// @router /login [post]
func (u *UserController) Login() {
	userLogin := &dtos.UserLoginDto{}
	json.Unmarshal(u.Ctx.Input.RequestBody, userLogin)
	if err := services.LoginUser(userLogin); err != nil {
		u.Data["json"] = map[string]string{"error": err.Error()}
	} else {
		message := fmt.Sprintf("Login success for user %v!", userLogin.Username)
		u.Data["json"] = map[string]string{"message": message}
	}

	u.ServeJSON()
}

// @Title Register
// @Description Register user into the system
// @Param 	body 	body 	dtos.UserDto 	true 	"Body of user register info"
// @Success 200 {string} message: "Register success for user: username"
// @Failure 403 {string} error: "message"
// @Accept json
// @router /register [post]
func (u *UserController) Register() {
	user := &dtos.UserDto{}
	json.Unmarshal(u.Ctx.Input.RequestBody, user)
	if err := services.RegisterUser(user); err != nil {
		u.Data["json"] = map[string]string{"error": err.Error()}
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
func (u *UserController) Logout() {
	u.Data["json"] = map[string]string{"message": "Logout success!"}
	u.ServeJSON()
}
