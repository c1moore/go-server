SHELL := /bin/bash

up:
	docker-compose up -d

down:
	docker-compose down --volumes

reset: up down

start: up
	go run ./cmd/go-service/