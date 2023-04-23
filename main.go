package main

import (
	_ "beego-car-rental/routers"

	settings "beego-car-rental/conf"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

func initialize() {
	settings.LoadSettings()

	logs.SetLogger(logs.AdapterConsole)
}

func main() {
	initialize()

	beego.Run()
}
