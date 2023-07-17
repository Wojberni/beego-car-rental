package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type CreateUserTable_20230717_212322 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateUserTable_20230717_212322{}
	m.Created = "20230717_212322"

	migration.Register("CreateUserTable_20230717_212322", m)
}

// Run the migrations
func (m *CreateUserTable_20230717_212322) Up() {
	m.SQL(`CREATE TABLE IF NOT EXISTS "user" (
        "id" serial NOT NULL PRIMARY KEY,
        "uuid" varchar(36) NOT NULL DEFAULT ''  UNIQUE,
        "username" varchar(32) NOT NULL DEFAULT '' ,
        "password" varchar(128) NOT NULL DEFAULT '' ,
        "email" varchar(64) NOT NULL DEFAULT ''  UNIQUE,
        "created" timestamp with time zone NOT NULL,
        "updated" timestamp with time zone NOT NULL
    );`)
}

// Reverse the migrations
func (m *CreateUserTable_20230717_212322) Down() {
	m.SQL(`DROP TABLE IF EXISTS "user" CASCADE;`)
}
