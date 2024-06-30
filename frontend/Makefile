SHELL := /bin/bash

build:
	npm run build
	
run:
	npm run dev

test:
	npm run test

docker-build:
	npm run build
	sudo docker-compose build --no-cache

docker-build-run:
	npm run build
	sudo docker-compose build --no-cache
	sudo docker-compose up

docker-publish:
	npm run build
	sudo docker-compose build --no-cache
	sudo docker tag task-management-frontend:1.0.0 ldilone25/task-management-frontend:1.0.0
	sudo docker push ldilone25/task-management-frontend:1.0.0