FROM golang:1.20.2 AS build-golang

ENV APP_NAME=beego-car-rental
WORKDIR /${APP_NAME}

COPY . .
RUN go get -v
RUN go build -v -o /usr/local/bin/beego-car-rental

EXPOSE 8080
CMD ["beego-car-rental"]
