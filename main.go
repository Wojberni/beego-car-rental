package main

import (
	_ "beego-car-rental/routers"

	settings "beego-car-rental/conf"

	beego "github.com/beego/beego/v2/server/web"
)

func initialize() {
	settings.LoadSettings()
}

func main() {
	initialize()

	beego.Run()
}
