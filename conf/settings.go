package settings

import (
	"fmt"
	"strconv"

	"github.com/beego/beego/v2/client/cache"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/config"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/lib/pq"
)

var (
	Profile    string
	AppVersion string
	DebugOrm   bool
	AppCache   cache.Cache
)

func LoadSettings() {
	logs.SetLogger(logs.AdapterConsole)

	cfg, err := config.NewConfig("ini", "conf/app.conf")
	if err != nil {
		panic(fmt.Sprintf("Failed to load config file!\nError message: %v", err))
	}

	Profile := cfg.DefaultString("runmode", "dev")
	AppVersion := cfg.DefaultString("appversion", "0.0.0")
	logs.Info("Using profile:", Profile)
	logs.Info("App version:", AppVersion)

	beego.BConfig.CopyRequestBody = cfg.DefaultBool("copyrequestbody", false)
	beego.BConfig.WebConfig.DirectoryIndex = cfg.DefaultBool("directoryindex", false)
	beego.BConfig.WebConfig.AutoRender = cfg.DefaultBool("autorender", false)
	beego.SetStaticPath("/", "static")

	beego.BConfig.WebConfig.Session.SessionOn = cfg.DefaultBool("sessionon", false)
	beego.BConfig.WebConfig.Session.SessionProvider = cfg.DefaultString("sessionprovider", "")
	beego.BConfig.WebConfig.Session.SessionName = cfg.DefaultString("sessionname", "")
	beego.BConfig.WebConfig.Session.SessionCookieLifeTime = cfg.DefaultInt("sessioncookielifetime", 0)
	beego.BConfig.WebConfig.Session.SessionGCMaxLifetime = cfg.DefaultInt64("sessiongcmaxlifetime", 0)

	cfgSection, err := cfg.GetSection(Profile)
	if err != nil {
		logs.Error(fmt.Printf("Invalid profile! You chose :%v, but available are [dev, test, prod]", Profile))
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

	if cfgSection["debugorm"] == "true" {
		DebugOrm = true
		orm.Debug = true
	} else {
		DebugOrm = false
		orm.Debug = false
	}

	err := orm.RegisterDriver("postgres", orm.DRPostgres)
	if err != nil {
		logs.Error("Error while registering orm driver:", err.Error())
	}
	err = orm.RegisterDataBase("default", "postgres", cfgSection["sqlconn"], orm.MaxOpenConnections(10), orm.MaxIdleConnections(5))
	if err != nil {
		logs.Error("Error while registering database:", err.Error())
	}
	orm.RunCommand()
	orm.RunSyncdb("default", false, DebugOrm)

	AppCache, err := cache.NewCache("memory", `{"interval":360}`)
	if err != nil {
		logs.Error("Error while creating new cache:", err.Error())
	} else {
		logs.Info("Cache creating successfull:", AppCache)
	}
}
