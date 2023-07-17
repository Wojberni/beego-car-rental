package controllers

import (
	"beego-car-rental/dtos"
	"beego-car-rental/models"
	"beego-car-rental/services"
	"encoding/json"
	"fmt"
)

// Operations about Roles
type RoleController struct {
	BaseController
}

// @Title GetAll
// @Description Get all Roles
// @Success 200 {object} models.Role
// @Failure 401 {string} error: "Unauthenticated, please log in!"
// @Failure 500 {string} error: "message"
// @Accept json
// @router / [get]
func (r *RoleController) GetAll() {
	roles := &models.RoleList{}
	if err := services.GetAllRoles(roles); err != nil {
		r.Data["json"] = map[string]string{"error": err.Error()}
	} else {
		r.Data["json"] = roles
	}
	r.ServeJSON()
}

// @Title Get
// @Description Get role by id
// @Param 	id 	path 	int 	true 	"The id of role to get"
// @Success 200 {object} models.Role
// @Failure 401 {string} error: "Unauthenticated, please log in!"
// @Failure 500 {string} error: "message"
// @Accept json
// @router /:id [get]
func (r *RoleController) Get() {
	id, _ := r.GetInt(":id")
	role := &models.Role{}
	if err := services.GetRole(role, id); err != nil {
		r.Data["json"] = map[string]string{"error": err.Error()}
	} else {
		r.Data["json"] = role
	}
	r.ServeJSON()
}

// @Title CreateRole
// @Description Create Role
// @Param	body		body 	dtos.RoleDto	true		"Body for Role content"
// @Success 201 {string} message: "Created role: Name"
// @Failure 401 {string} error: "Unauthenticated, please log in!"
// @Failure 500 {string} error: "message"
// @Accept json
// @router / [post]
func (u *RoleController) Post() {
	role := &dtos.RoleDto{}
	json.Unmarshal(u.Ctx.Input.RequestBody, role)
	if err := services.CreateRole(role); err != nil {
		u.Data["json"] = map[string]string{"error": err.Error()}
	} else {
		message := fmt.Sprintf("Created role: %v!", role.Name)
		u.Data["json"] = map[string]string{"message": message}
	}
	u.ServeJSON()
}

// @Title Update
// @Description Update the role
// @Param 	id 	path 	int 			true 	"The id you want to update"
// @Param	body 	body 	dtos.RoleDto 	true 	"Body for role content"
// @Success 200 {string} message: "Updated role: id"
// @Failure 401 {string} error: "Unauthenticated, please log in!"
// @Failure 500 {string} error: "message"
// @Accept json
// @router /:id [put]
func (u *RoleController) Put() {
	id, _ := u.GetInt(":id")
	role := &dtos.RoleDto{}
	json.Unmarshal(u.Ctx.Input.RequestBody, role)
	if err := services.UpdateRole(role, id); err != nil {
		u.Data["json"] = map[string]string{"error": err.Error()}
	} else {
		message := fmt.Sprintf("Updated role: %v", id)
		u.Data["json"] = map[string]string{"message": message}
	}
	u.ServeJSON()
}

// @Title Delete
// @Description delete the role
// @Param	id		path 	int	true		"The id you want to delete"
// @Success 200 {string} message: "Deleted role: id"
// @Failure 401 {string} error: "Unauthenticated, please log in!"
// @Failure 500 {string} error: "message"
// @Accept json
// @router /:id [delete]
func (u *RoleController) Delete() {
	id, _ := u.GetInt(":id")
	if err := services.DeleteRole(id); err != nil {
		u.Data["json"] = map[string]string{"error": err.Error()}
	} else {
		message := fmt.Sprintf("Deleted role: %v", id)
		u.Data["json"] = map[string]string{"message": message}
	}
	u.ServeJSON()
}
