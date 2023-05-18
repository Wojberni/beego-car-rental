package models

import (
	"time"

	"github.com/beego/beego/v2/adapter/orm"
)

func init() {
	orm.RegisterModel(new(Car))
}

type CarList []Car

type Car struct {
	Id       int
	Uuid     string    `orm:"size(32);unique"`
	Make     string    `orm:"size(32)"`
	Model    string    `orm:"size(32)"`
	RegPlate string    `orm:"size(12);unique"`
	Created  time.Time `orm:"auto_now_add"`
	Updated  time.Time `orm:"auto_now"`
}

func (u *Car) Insert() error {
	if _, err := orm.NewOrm().Insert(u); err != nil {
		return err
	}
	return nil
}

func (u *Car) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(u, fields...); err != nil {
		return err
	}
	return nil
}

func (u *Car) Delete() error {
	if _, err := orm.NewOrm().Delete(u); err != nil {
		return err
	}
	return nil
}

func (u *Car) Read(fields ...string) error {
	if err := orm.NewOrm().Read(u, fields...); err != nil {
		return err
	}
	return nil
}

func (u *CarList) ReadAll() error {
	qb, _ := orm.NewQueryBuilder("postgres")
	qb.Select("*").From("user").Limit(100).Offset(0)
	sql := qb.String()
	if _, err := orm.NewOrm().Raw(sql).QueryRows(u); err != nil {
		return err
	}
	return nil
}
