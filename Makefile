DB_DEV_URL=postgres://dev:1234@localhost/beego_car_rental_dev?sslmode=disable

run : 
	bee run -gendoc -downdoc

docker-run :
	docker-compose up -d

docker-stop :
	docker-compose down -d

db-migrate-dev :
	bee migrate -driver=postgres -conn=$(DB_DEV_URL)

db-rollback-dev :
	bee migrate rollback -driver=postgres -conn=$(DB_DEV_URL)

db-reset-dev : 
	bee migrate reset -driver=postgres -conn=$(DB_DEV_URL)

db-update-dev : 
	bee migrate refresh -driver=postgres -conn=$(DB_DEV_URL)


.PHONY : run docker-run docker-stop db-migrate-dev db-rollback-dev db-reset-dev db-update-dev