package settings

import (
	"fmt"
	"strconv"

	"github.com/beego/beego/v2/core/config"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

func LoadSettings() {

	cfg, err := config.NewConfig("ini", "conf/app.conf")
	if err != nil {
		panic(fmt.Sprintf("Failed to load config file!\nError message: %v", err))
	}

	profile := cfg.DefaultString("runmode", "dev")
	logs.Info("Using profile:", profile)

	beego.BConfig.CopyRequestBody = cfg.DefaultBool("copyrequestbody", false)
	beego.BConfig.WebConfig.DirectoryIndex = cfg.DefaultBool("directoryindex", false)
	beego.BConfig.WebConfig.AutoRender = cfg.DefaultBool("autorender", true)
	beego.SetStaticPath("/", "views/static")

	cfgSection, err := cfg.GetSection(profile)
	if err != nil {
		logs.Error(fmt.Printf("Invalid profile! You chose :%v, but available are [dev, test, prod]", profile))
	}

	loadProfileSettings(cfgSection)
}

func loadProfileSettings(cfgSection map[string]string) {
	if cfgSection["enabledocs"] == "true" {
		beego.SetStaticPath("/swagger", "swagger")
	} else {
		beego.DelStaticPath("/swagger")
	}

	port, _ := strconv.Atoi(cfgSection["httpport"])
	beego.BConfig.Listen.HTTPPort = port

	if cfgSection["recoverpanic"] == "true" {
		beego.BConfig.RecoverPanic = true
	} else {
		beego.BConfig.RecoverPanic = false
	}

	if cfgSection["enableadmin"] == "true" {
		beego.BConfig.Listen.EnableAdmin = true
		beego.BConfig.Listen.AdminAddr = "localhost"
		beego.BConfig.Listen.AdminPort = 8088
	} else {
		beego.BConfig.Listen.EnableAdmin = false
	}

	// todo: register database here with orm
}
