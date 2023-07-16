package models

import (
	"fmt"
	"time"

	"github.com/beego/beego/v2/adapter/orm"
	"github.com/beego/beego/v2/core/logs"
)

func init() {
	orm.RegisterModel(new(Role))
}

type RoleList []Role

type Role struct {
	Id         int          `orm:"auto;pk;column(id)"`
	Name       string       `orm:"size(32)"`
	Users      []*User      `orm:"reverse(many);rel_table(user_role)"`
	Privileges []*Privilege `orm:"rel(m2m);rel_table(role_privilege)"`
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
	// todo: fix select for role
	qb.Select("*").
		From("role").
		InnerJoin("role_privilege").On("role.id = role_privilege.role_id").
		InnerJoin("privilege").On("privilege.id = role_privilege.privilege_id").
		Limit(100).
		Offset(0)
	sql := qb.String()
	if _, err := orm.NewOrm().Raw(sql).QueryRows(r); err != nil {
		return err
	}
	logs.Info(fmt.Sprintf("%v", r))
	return nil
}
