package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type CreateRoleTable_20230717_212604 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateRoleTable_20230717_212604{}
	m.Created = "20230717_212604"

	migration.Register("CreateRoleTable_20230717_212604", m)
}

// Run the migrations
func (m *CreateRoleTable_20230717_212604) Up() {
	m.SQL(`CREATE TABLE IF NOT EXISTS "role" (
        "id" serial NOT NULL PRIMARY KEY,
        "name" varchar(32) NOT NULL DEFAULT '' ,
        "created" timestamp with time zone NOT NULL,
        "updated" timestamp with time zone NOT NULL
    );`)
}

// Reverse the migrations
func (m *CreateRoleTable_20230717_212604) Down() {
	m.SQL(`DROP TABLE IF EXISTS "role" CASCADE;`)
}
