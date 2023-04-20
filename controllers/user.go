package controllers

import (
	"encoding/json"
	"fmt"
	"go-car-rental/models"

	beego "github.com/beego/beego/v2/server/web"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

// @APIVersion 1.0.0
// @Title CreateUser
// @Description Create user
// @Param 	body 	body 	models.User 	true 	"Body for user content"
// @Success 200 {string} Created user with Uuid!
// @Failure 403 {string} Data missing, please fill all data!
// @Accept json
// @router / [post]
func (u *UserController) Post() {
	var user models.User
	var message string
	json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	// todo: move validation inside model
	if user.Email == "" || user.Password == "" || user.Username == "" {
		message = "Data missing, please fill all data!"
	} else {
		uid := models.AddUser(user)
		message = fmt.Sprintf("Created user with Uuid %s", uid)
	}

	u.Data["json"] = map[string]string{"message": message}
	u.ServeJSON()
}

// @APIVersion 1.0.0
// @Title GetAll
// @Description Get all Users
// @Success 200 {object} models.User
// @Failure 404 {string} Error retrieving data, please try again later!
// @Accept json
// @router / [get]
func (u *UserController) GetAll() {
	users := models.GetAllUsers()
	u.Data["json"] = users
	u.ServeJSON()
}

// @APIVersion 1.0.0
// @Title Get
// @Description Get user by uuid
// @Param 	uid 	path 	string 	true 	"The uuid of user to get"
// @Success 200 {object} models.User
// @Failure 403 {string} Uuid is empty!
// @Accept json
// @router /:uuid [get]
func (u *UserController) Get() {
	uid := u.GetString(":uuid")
	if uid != "" {
		user, err := models.GetUser(uid)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = user
		}
	}
	u.ServeJSON()
}

// @APIVersion 1.0.0
// @Title Update
// @Description Update the user
// @Param 	uuid 	path 	string 			true 	"The uuid you want to update"
// @Param	body 	body 	models.User 	true 	"Body for user content"
// @Success 200 {object} models.User
// @Failure 403 {string} Uuid is not int!
// @Accept json
// @router /:uid [put]
func (u *UserController) Put() {
	uid := u.GetString(":uid")
	if uid != "" {
		var user models.User
		json.Unmarshal(u.Ctx.Input.RequestBody, &user)
		uu, err := models.UpdateUser(uid, &user)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = uu
		}
	}
	u.ServeJSON()
}

// @APIVersion 1.0.0
// @Title Delete
// @Description delete the user
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} Deleted user: uuid
// @Failure 403 {string} User not found: uuid
// @Accept json
// @router /:uid [delete]
func (u *UserController) Delete() {
	uid := u.GetString(":uid")
	var message string
	if models.DeleteUser(uid) {
		message = fmt.Sprintf("Deleted user: %s", uid)
	} else {
		message = fmt.Sprintf("User not found: %s", uid)
	}
	u.Data["json"] = message
	u.ServeJSON()
}

// @APIVersion 1.0.0
// @Title Login
// @Description Logs user into the system
// @Param	username		query 	string	true		"The username for login"
// @Param	password		query 	string	true		"The password for login"
// @Success 200 {string} Login success!
// @Failure 403 {string} User does not exist!
// @Accept json
// @router /login [get]
func (u *UserController) Login() {
	username := u.GetString("username")
	password := u.GetString("password")
	if models.Login(username, password) {
		u.Data["json"] = "Login success!"
	} else {
		u.Data["json"] = "User does not exist!"
	}
	u.ServeJSON()
}

// @APIVersion 1.0.0
// @Title Register
// @Description Register user into the system
// @Param	username		query 	string	true		"The username for register"
// @Param	password		query 	string	true		"The password for register"
// @Param	email		query 	string	true		"The email for register"
// @Success 200 {string} Register success!
// @Failure 403 {string} Register failure! Fill all fields!
// @Accept json
// @router /register [get]
func (u *UserController) Register() {
	username := u.GetString("username")
	password := u.GetString("password")
	email := u.GetString("email")
	if models.Register(username, password, email) {
		u.Data["json"] = "Register success!"
	} else {
		u.Data["json"] = "Register failure! Fill all fields!"
	}
	u.ServeJSON()
}

// @APIVersion 1.0.0
// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} Logout success!
// @Accept json
// @router /logout [get]
func (u *UserController) Logout() {
	u.Data["json"] = "Logout success!"
	u.ServeJSON()
}
