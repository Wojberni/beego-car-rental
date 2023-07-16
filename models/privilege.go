package models

import (
	"time"

	"github.com/beego/beego/v2/adapter/orm"
)

func init() {
	orm.RegisterModel(new(Privilege))
}

type PrivilegeList []Privilege

type Privilege struct {
	Id      int       `orm:"auto;pk;column(id)"`
	Name    string    `orm:"size(64)"`
	Roles   []*Role   `orm:"reverse(many);rel_table(role_privilege)"`
	Created time.Time `orm:"auto_now_add"`
	Updated time.Time `orm:"auto_now"`
}

func (p *Privilege) Insert() error {
	if _, err := orm.NewOrm().Insert(p); err != nil {
		return err
	}
	return nil
}

func (p *Privilege) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(p, fields...); err != nil {
		return err
	}
	return nil
}

func (p *Privilege) Delete() error {
	if _, err := orm.NewOrm().Delete(p); err != nil {
		return err
	}
	return nil
}

func (p *Privilege) Read(fields ...string) error {
	if err := orm.NewOrm().Read(p, fields...); err != nil {
		return err
	}
	return nil
}

func (p *PrivilegeList) ReadAll() error {
	qb, _ := orm.NewQueryBuilder("postgres")
	qb.Select("*").From("privilege").Limit(100).Offset(0)
	sql := qb.String()
	if _, err := orm.NewOrm().Raw(sql).QueryRows(p); err != nil {
		return err
	}
	return nil
}
