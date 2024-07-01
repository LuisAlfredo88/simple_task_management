SHELL := /bin/bash

build:
	go build -o ./build/main ./cmd/bin/
	
run:
	source .env	
	go run cmd/bin/main.go

test:
	source .env
	go test ./...

migrate:
	go run cmd/migration/main.go

docker-build:
	go build -o ./build/main ./cmd/bin/
	sudo docker-compose build --no-cache

docker-buil-run:
	go build -o ./build/main ./cmd/bin/
	sudo docker-compose build --no-cache
	sudo docker-compose up

docker-publish:
	go build -o ./build/main ./cmd/bin/
	sudo docker-compose build --no-cache
	sudo docker tag task-management-backend:1.0.0 ldilone25/task-management-backend:1.0.0
	sudo docker push ldilone25/task-management-backend:1.0.0