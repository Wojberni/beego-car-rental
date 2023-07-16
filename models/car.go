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
	Id       int       `orm:"auto;pk;column(id)"`
	Make     string    `orm:"size(32)"`
	Model    string    `orm:"size(32)"`
	RegPlate string    `orm:"size(12);unique"`
	Orders   []*Order  `orm:"reverse(many)"`
	Created  time.Time `orm:"auto_now_add"`
	Updated  time.Time `orm:"auto_now"`
}

func (c *Car) Insert() error {
	if _, err := orm.NewOrm().Insert(c); err != nil {
		return err
	}
	return nil
}

func (c *Car) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(c, fields...); err != nil {
		return err
	}
	return nil
}

func (c *Car) Delete() error {
	if _, err := orm.NewOrm().Delete(c); err != nil {
		return err
	}
	return nil
}

func (c *Car) Read(fields ...string) error {
	if err := orm.NewOrm().Read(c, fields...); err != nil {
		return err
	}
	return nil
}

func (c *CarList) ReadAll() error {
	qb, _ := orm.NewQueryBuilder("postgres")
	qb.Select("*").From("car").Limit(100).Offset(0)
	sql := qb.String()
	if _, err := orm.NewOrm().Raw(sql).QueryRows(c); err != nil {
		return err
	}
	return nil
}
