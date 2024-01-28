.PHONY: default run build test docs clean

APP_NAME=gopportunities

default: run

run:
	@go run main.go

build:
	@go build -o $(APP_NAME) main.go
test:
	@go test ./ ...
dics:
	@swag init
