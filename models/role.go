package models

import (
	"time"

	"github.com/beego/beego/v2/adapter/orm"
)

func init() {
	orm.RegisterModel(new(Role))
}

type RoleList []Role

// relation many to many to user, many to many to privilege

type Role struct {
	Id         int
	Name       string       `orm:"size(32)"`
	Users      []*User      `orm:"reverse(many)"`
	Privileges []*Privilege `orm:"rel(m2m)"`
	Created    time.Time    `orm:"auto_now_add"`
	Updated    time.Time    `orm:"auto_now"`
}

func (r *Role) Insert() error {
	if _, err := orm.NewOrm().Insert(r); err != nil {
		return err
	}
	return nil
}

func (r *Role) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(r, fields...); err != nil {
		return err
	}
	return nil
}

func (r *Role) Delete() error {
	if _, err := orm.NewOrm().Delete(r); err != nil {
		return err
	}
	return nil
}

func (r *Role) Read(fields ...string) error {
	if err := orm.NewOrm().Read(r, fields...); err != nil {
		return err
	}
	return nil
}

func (r *RoleList) ReadAll() error {
	qb, _ := orm.NewQueryBuilder("postgres")
	qb.Select("*").From("role").Limit(100).Offset(0)
	sql := qb.String()
	if _, err := orm.NewOrm().Raw(sql).QueryRows(r); err != nil {
		return err
	}
	return nil
}
