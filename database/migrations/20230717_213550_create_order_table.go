package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type CreateOrderTable_20230717_213550 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateOrderTable_20230717_213550{}
	m.Created = "20230717_213550"

	migration.Register("CreateOrderTable_20230717_213550", m)
}

// Run the migrations
func (m *CreateOrderTable_20230717_213550) Up() {
	m.SQL(`CREATE TABLE IF NOT EXISTS "order" (
        "id" serial NOT NULL PRIMARY KEY,
        "car_id" integer NOT NULL REFERENCES "car" ("id"),
        "user_id" integer NOT NULL REFERENCES "user" ("id"),
        "created" timestamp with time zone NOT NULL,
        "updated" timestamp with time zone NOT NULL
    );`)
}

// Reverse the migrations
func (m *CreateOrderTable_20230717_213550) Down() {
	m.SQL(`DROP TABLE IF EXISTS "order" CASCADE;`)
}
