package models

import (
	"time"

	"github.com/beego/beego/v2/adapter/orm"
)

func init() {
	orm.RegisterModel(new(Order))
}

type OrderList []Order

// relation one to many to user, one to many to car

type Order struct {
	Id      int
	Cars    *Car      `orm:"rel(fk)"`
	Users   *User     `orm:"rel(fk)"`
	Created time.Time `orm:"auto_now_add"`
	Updated time.Time `orm:"auto_now"`
}

func (o *Order) Insert() error {
	if _, err := orm.NewOrm().Insert(o); err != nil {
		return err
	}
	return nil
}

func (o *Order) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(o, fields...); err != nil {
		return err
	}
	return nil
}

func (o *Order) Delete() error {
	if _, err := orm.NewOrm().Delete(o); err != nil {
		return err
	}
	return nil
}

func (o *Order) Read(fields ...string) error {
	if err := orm.NewOrm().Read(o, fields...); err != nil {
		return err
	}
	return nil
}

func (o *OrderList) ReadAll() error {
	qb, _ := orm.NewQueryBuilder("postgres")
	qb.Select("*").From("order").Limit(100).Offset(0)
	sql := qb.String()
	if _, err := orm.NewOrm().Raw(sql).QueryRows(o); err != nil {
		return err
	}
	return nil
}
