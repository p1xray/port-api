.PHONY: dc build run test lint

dc:
	docker-compose up --remove-orphans -- build

build:
	go build -v -o app.exe cmd/port-api/main.go

run:
	go run cmd/port-api/main.go

test:
	go test -race ./..

lint:
	golangci-lint run