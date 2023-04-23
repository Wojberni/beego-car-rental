# Build Golang binary
FROM golang:1.20.2 AS build-golang

WORKDIR /home/wojciech/GoProjects/beego-car-rental

COPY . .
RUN go get -v && go build -v -o /usr/local/bin/beego-car-rental

EXPOSE 8080
CMD ["beego-car-rental"]
