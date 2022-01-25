VERSION = $(shell git branch --show-current)

DB_CONNECTION = "user=docker password=docker dbname=go-clean-arch sslmode=disable"

mig-create: 
	goose -dir ./migrations create $(MIG_NAME) sql 

mig-status:
	goose postgres $(DB_CONNECTION) status

mig-up: 
	goose -dir ./migrations postgres $(DB_CONNECTION) up

mig-down: 
	goose -dir ./migrations postgres $(DB_CONNECTION) down
