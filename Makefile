# app name
APP_NAME := server	

dev:
	go run ./cmd/${APP_NAME}

run: 
	docker compose up -d && go run ./cmd/${APP_NAME}

kill:
	docker compose kill

up:
	docker compose up -d

down:
	docker compose down

.PHONY: run

.PHONY: air