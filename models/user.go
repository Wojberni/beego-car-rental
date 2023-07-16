package models

import (
	"time"

	"github.com/beego/beego/v2/adapter/orm"
)

func init() {
	orm.RegisterModel(new(User))
}

type UserList []User

type User struct {
	Id       int       `orm:"auto;pk;column(id)"`
	Uuid     string    `orm:"size(36);unique"`
	Username string    `orm:"size(32)"`
	Password string    `orm:"size(128)"`
	Email    string    `orm:"size(64);unique"`
	Roles    []*Role   `orm:"rel(m2m);rel_table(user_role)"`
	Orders   []*Order  `orm:"reverse(many)"`
	Created  time.Time `orm:"auto_now_add"`
	Updated  time.Time `orm:"auto_now"`
}

func (u *User) Insert() error {
	if _, err := orm.NewOrm().Insert(u); err != nil {
		return err
	}
	return nil
}

func (u *User) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(u, fields...); err != nil {
		return err
	}
	return nil
}

func (u *User) Delete() error {
	if _, err := orm.NewOrm().Delete(u); err != nil {
		return err
	}
	return nil
}

func (u *User) Read(fields ...string) error {
	if err := orm.NewOrm().Read(u, fields...); err != nil {
		return err
	}
	return nil
}

func (u *UserList) ReadAll() error {
	qb, _ := orm.NewQueryBuilder("postgres")
	qb.Select("*").From("user").Limit(100).Offset(0)
	sql := qb.String()
	if _, err := orm.NewOrm().Raw(sql).QueryRows(u); err != nil {
		return err
	}
	return nil
}
