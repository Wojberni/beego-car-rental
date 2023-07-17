FROM golang:1.20.2 AS migrate-golang

ENV APP_NAME=beego-car-rental-migrations
WORKDIR /${APP_NAME}

# move migrations to dockerfile, so database can start

COPY database/ .
RUN go install github.com/beego/bee/v2@latest
RUN bee migrate -driver=postgres -conn=postgres://dev:1234@192.168.0.10/beego_car_rental_dev?sslmode=disable -dir=migrations

FROM golang:1.20.2 AS build-golang

ENV APP_NAME=beego-car-rental
WORKDIR /${APP_NAME}

COPY . .
RUN go get -v

RUN go build -v -o /usr/local/bin/beego-car-rental

EXPOSE 8080
CMD ["beego-car-rental"]
