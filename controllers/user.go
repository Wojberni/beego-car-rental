package controllers

import (
	"beego-car-rental/dtos"
	"beego-car-rental/models"
	"beego-car-rental/services"
	"encoding/json"
	"fmt"
)

// Operations about Users
type UserController struct {
	BaseController
}

// @Title GetAll
// @Description Get all Users
// @Success 200 {object} models.User
// @Failure 401 {string} error: "Unauthenticated, please log in!"
// @Failure 403 {string} error: "Unauthorized, forbidden content access!"
// @Failure 500 {string} error: "message"
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
// @Failure 401 {string} error: "Unauthenticated, please log in!"
// @Failure 403 {string} error: "Unauthorized, forbidden content access!"
// @Failure 500 {string} error: "message"
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

// @Title CreateUser
// @Description Create User
// @Param	body		body 	dtos.UserDto	true		"Body for User content"
// @Success 201 {string} message: "Created user: Username"
// @Failure 401 {string} error: "Unauthenticated, please log in!"
// @Failure 403 {string} error: "Unauthorized, forbidden content access!"
// @Failure 500 {string} error: "message"
// @Accept json
// @router / [post]
func (u *UserController) Post() {
	user := &dtos.UserDto{}
	json.Unmarshal(u.Ctx.Input.RequestBody, user)
	if err := services.CreateUser(user); err != nil {
		u.Data["json"] = map[string]string{"error": err.Error()}
	} else {
		message := fmt.Sprintf("Created user: %v!", user.Username)
		u.Data["json"] = map[string]string{"message": message}
	}
	u.ServeJSON()
}

// @Title Update
// @Description Update the user
// @Param 	id 	path 	int 			true 	"The id you want to update"
// @Param	body 	body 	dtos.UserDto 	true 	"Body for user content"
// @Success 200 {string} message: "Updated user: id"
// @Failure 401 {string} error: "Unauthenticated, please log in!"
// @Failure 403 {string} error: "Unauthorized, forbidden content access!"
// @Failure 500 {string} error: "message"
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
// @Failure 401 {string} error: "Unauthenticated, please log in!"
// @Failure 403 {string} error: "Unauthorized, forbidden content access!"
// @Failure 500 {string} error: "message"
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
