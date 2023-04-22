# Build Golang binary
FROM golang:1.20.2 AS build-golang

WORKDIR /home/wojciech/GoProjects/go-car-rental

COPY . .
RUN go get -v && go build -v -o /usr/local/bin/go-car-rental

EXPOSE 8080
CMD ["go-car-rental"]
