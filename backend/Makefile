SHELL := /bin/bash

build:
	go build -o cmd/bin/main.go

run:
	source .env	
	go run cmd/bin/main.go

test:
	source .env
	go test ./...

migrate:
	go run cmd/migration/main.go

air:
	source .env	
	air cmd/bin/main.go