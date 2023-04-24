package main

import (
	_ "beego-car-rental/routers"

	settings "beego-car-rental/conf"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/lib/pq"
)

type TestUser struct {
	ID   int    `orm:"column(id)"`
	Name string `orm:"column(name)"`
}

func initialize() {
	settings.LoadSettings()

	orm.Debug = true

	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default", "postgres", "postgres://dev:1234@localhost/beego_car_rental_dev?sslmode=disable")

	logs.SetLogger(logs.AdapterConsole)
}

func main() {
	initialize()

	orm.RegisterModel(new(TestUser))
	orm.RunSyncdb("default", false, true)

	// create orm object
	o := orm.NewOrm()

	// data
	user := new(TestUser)
	user.Name = "Mike"
	user.ID = 0

	// insert data
	o.Insert(user)

	beego.Run()
}
