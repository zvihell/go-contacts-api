build:
	docker-compose build 
run:
	docker-compose up 

migrate:
	migrate -path ./schema -database 'postgres://postgres:qwerty@0.0.0.0:5433/postgres?sslmode=disable' up
