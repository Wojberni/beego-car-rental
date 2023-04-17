package main

import (
	_ "go-car-rental/routers"

	"github.com/beego/beego/v2/core/config"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {

	logs.SetLogger(logs.AdapterConsole)

	cfg, err := config.NewConfig("ini", "conf/dev.ini")
	if err != nil {
		logs.Error(err)
	}

	profile, _ := cfg.String("runmode")
	logs.Info("Using profile:", profile)

	beego.SetStaticPath("/", "views/static")

	if cfg.DefaultBool("EnableDocs", false) {
		beego.SetStaticPath("/swagger", "/swagger")
	}

	if cfg.DefaultBool("EnableAdmin", false) {
		beego.BConfig.Listen.EnableAdmin = true
		beego.BConfig.Listen.AdminAddr = "localhost"
		beego.BConfig.Listen.AdminPort = 8088
	}

	beego.Run()
}
