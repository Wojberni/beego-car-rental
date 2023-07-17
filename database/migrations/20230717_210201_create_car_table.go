package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type CreateCarTable_20230717_210201 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateCarTable_20230717_210201{}
	m.Created = "20230717_210201"

	migration.Register("CreateCarTable_20230717_210201", m)
}

// Run the migrations
func (m *CreateCarTable_20230717_210201) Up() {
	m.SQL(`CREATE TABLE IF NOT EXISTS "car" (
        "id" serial NOT NULL PRIMARY KEY,
        "make" varchar(32) NOT NULL DEFAULT '' ,
        "model" varchar(32) NOT NULL DEFAULT '' ,
        "reg_plate" varchar(12) NOT NULL DEFAULT ''  UNIQUE,
        "created" timestamp with time zone NOT NULL,
        "updated" timestamp with time zone NOT NULL);`)
}

// Reverse the migrations
func (m *CreateCarTable_20230717_210201) Down() {
	m.SQL(`DROP TABLE IF EXISTS "car" CASCADE;`)
}
