#!/bin/bash

export DB_NAME_DEV="BEEGO_CAR_RENTAL_DEV"
export DB_NAME_TEST="BEEGO_CAR_RENTAL_TEST"
export DB_NAME_PROD="BEEGO_CAR_RENTAL_PROD"

export DEV_DB_USER="dev"
export DEV_DB_PASSWD="1234"

export TEST_DB_USER="test"
export TEST_DB_PASSWD="1234"

export PROD_DB_USER="prod"
export PROD_DB_PASSWD="1234"

cd /init/sql

psql -U postgres -v db=$DB_NAME_DEV -f create_database.sql
psql -U postgres -v db=$DB_NAME_TEST -f create_database.sql
psql -U postgres -v db=$DB_NAME_PROD -f create_database.sql

psql -U postgres -v db=$DB_NAME_DEV -v user=$DEV_DB_USER -v passwd="'${DEV_DB_PASSWD}'" -f create_user.sql
psql -U postgres -v db=$DB_NAME_TEST -v user=$TEST_DB_USER -v passwd="'${TEST_DB_PASSWD}'" -f create_user.sql
psql -U postgres -v db=$DB_NAME_PROD -v user=$PROD_DB_USER -v passwd="'${PROD_DB_PASSWD}'" -f create_user.sql
