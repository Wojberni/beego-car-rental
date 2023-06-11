package controllers

import (
	"beego-car-rental/dtos"
	"beego-car-rental/models"
	"beego-car-rental/services"
	"encoding/json"
	"fmt"
)

// Operations about Users
type PrivilegeController struct {
	BaseController
}

// @Title GetAll
// @Description Get all Privileges
// @Success 200 {object} models.Privilege
// @Failure 404 {string} error: "message"
// @Accept json
// @router / [get]
func (p *PrivilegeController) GetAll() {
	privileges := &models.PrivilegeList{}
	if err := services.GetAllPrivileges(privileges); err != nil {
		p.Data["json"] = map[string]string{"error": err.Error()}
	} else {
		p.Data["json"] = privileges
	}
	p.ServeJSON()
}

// @Title Get
// @Description Get privilege by id
// @Param 	id 	path 	int 	true 	"The id of privilege to get"
// @Success 200 {object} models.Privilege
// @Failure 403 {string} error: "message"
// @Accept json
// @router /:id [get]
func (p *PrivilegeController) Get() {
	id, _ := p.GetInt(":id")
	privilege := &models.Privilege{}
	if err := services.GetPrivilege(privilege, id); err != nil {
		p.Data["json"] = map[string]string{"error": err.Error()}
	} else {
		p.Data["json"] = privilege
	}
	p.ServeJSON()
}

// @Title CreatePrivilege
// @Description Create Privilege
// @Param	body		body 	dtos.PrivilegeDto	true		"Body for Privilege content"
// @Success 201 {string} message: "Created privilege: Name"
// @Failure 403 {string} error: "message"
// @Accept json
// @router / [post]
func (p *PrivilegeController) Post() {
	privilege := &dtos.PrivilegeDto{}
	json.Unmarshal(p.Ctx.Input.RequestBody, privilege)
	if err := services.CreatePrivilege(privilege); err != nil {
		p.Data["json"] = map[string]string{"error": err.Error()}
	} else {
		message := fmt.Sprintf("Created privilege: %v!", privilege.Name)
		p.Data["json"] = map[string]string{"message": message}
	}
	p.ServeJSON()
}

// @Title Update
// @Description Update the privilege
// @Param 	id 	path 	int 			true 	"The id you want to update"
// @Param	body 	body 	dtos.PrivilegeDto 	true 	"Body for privilege content"
// @Success 200 {string} message: "Updated privilege: id"
// @Failure 403 {string} error: "message"
// @Accept json
// @router /:id [put]
func (p *PrivilegeController) Put() {
	id, _ := p.GetInt(":id")
	privilege := &dtos.PrivilegeDto{}
	json.Unmarshal(p.Ctx.Input.RequestBody, privilege)
	if err := services.UpdatePrivilege(privilege, id); err != nil {
		p.Data["json"] = map[string]string{"error": err.Error()}
	} else {
		message := fmt.Sprintf("Updated privilege: %v", id)
		p.Data["json"] = map[string]string{"message": message}
	}
	p.ServeJSON()
}

// @Title Delete
// @Description delete the privilege
// @Param	id		path 	int	true		"The id you want to delete"
// @Success 200 {string} message: "Deleted privilege: id"
// @Failure 403 {string} error: "message"
// @Accept json
// @router /:id [delete]
func (p *PrivilegeController) Delete() {
	id, _ := p.GetInt(":id")
	if err := services.DeletePrivilege(id); err != nil {
		p.Data["json"] = map[string]string{"error": err.Error()}
	} else {
		message := fmt.Sprintf("Deleted privilege: %v", id)
		p.Data["json"] = map[string]string{"message": message}
	}
	p.ServeJSON()
}
