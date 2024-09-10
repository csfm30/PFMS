version := 1.1.5

doc:
	swag init

database_start:
	docker-compose up -d

database_down:
	docker-compose down



