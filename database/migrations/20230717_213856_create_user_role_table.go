package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type CreateUserRoleTable_20230717_213856 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateUserRoleTable_20230717_213856{}
	m.Created = "20230717_213856"

	migration.Register("CreateUserRoleTable_20230717_213856", m)
}

// Run the migrations
func (m *CreateUserRoleTable_20230717_213856) Up() {
	m.SQL(`CREATE TABLE IF NOT EXISTS "user_role" (
        "id" serial NOT NULL PRIMARY KEY,
        "user_id" integer NOT NULL REFERENCES "user" ("id"),
        "role_id" integer NOT NULL REFERENCES "role" ("id")
    );`)
}

// Reverse the migrations
func (m *CreateUserRoleTable_20230717_213856) Down() {
	m.SQL(`DROP TABLE IF EXISTS "user_role" CASCADE;`)
}
