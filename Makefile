version := 1.1.3

doc:
	swag init

database_start:
	docker-compose up -d

database_down:
	docker-compose down



