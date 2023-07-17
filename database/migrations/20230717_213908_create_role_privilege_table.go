package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type CreateRolePrivilegeTable_20230717_213908 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateRolePrivilegeTable_20230717_213908{}
	m.Created = "20230717_213908"

	migration.Register("CreateRolePrivilegeTable_20230717_213908", m)
}

// Run the migrations
func (m *CreateRolePrivilegeTable_20230717_213908) Up() {
	m.SQL(`CREATE TABLE IF NOT EXISTS "role_privilege" (
        "id" serial NOT NULL PRIMARY KEY,
        "role_id" integer NOT NULL REFERENCES "role" ("id"),
        "privilege_id" integer NOT NULL REFERENCES "privilege" ("id")
    );`)

}

// Reverse the migrations
func (m *CreateRolePrivilegeTable_20230717_213908) Down() {
	m.SQL(`DROP TABLE IF EXISTS "role_privilege" CASCADE;`)
}
