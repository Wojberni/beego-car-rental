package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type CreatePrivilegeTable_20230717_212016 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreatePrivilegeTable_20230717_212016{}
	m.Created = "20230717_212016"

	migration.Register("CreatePrivilegeTable_20230717_212016", m)
}

// Run the migrations
func (m *CreatePrivilegeTable_20230717_212016) Up() {
	m.SQL(`CREATE TABLE IF NOT EXISTS "privilege" (
        "id" serial NOT NULL PRIMARY KEY,
        "name" varchar(64) NOT NULL DEFAULT '' ,
        "created" timestamp with time zone NOT NULL,
        "updated" timestamp with time zone NOT NULL);`)
}

// Reverse the migrations
func (m *CreatePrivilegeTable_20230717_212016) Down() {
	m.SQL(`DROP TABLE IF EXISTS "privilege" CASCADE;`)
}
