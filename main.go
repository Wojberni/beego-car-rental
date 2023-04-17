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

	if profile == "dev" || profile == "test" {
		beego.SetStaticPath("/swagger", "/swagger")
		beego.SetStaticPath("/", "views/static")
	}

	beego.Run()
}
