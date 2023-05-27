package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

// Base controller with session authentication
type BaseController struct {
	beego.Controller
}

// Prepare function executes before each request
func (b *BaseController) Prepare() {
	session := b.GetSession("login")
	if session == nil {
		b.Data["json"] = map[string]string{"error": "Unauthenticated, please log in!"}
		b.ServeJSON()
		return
	}
}
