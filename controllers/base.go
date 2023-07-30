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
		b.Ctx.Output.SetStatus(401)
		b.Data["json"] = map[string]string{"error": "Unauthenticated, please log in!"}
		b.ServeJSON()
		return
	}
	if params, ok := session.([]interface{}); ok {
		// todo: check user privilege?
		if params[2] != "USER" {
			b.Ctx.Output.SetStatus(403)
			b.Data["json"] = map[string]string{"error": "Unauthorized, forbidden content access!"}
			b.ServeJSON()
			return
		}
	}
}
