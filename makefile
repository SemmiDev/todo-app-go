# run backend service
run:
	@go build -v -o bin/todoapp backend/main.go
	@./bin/todoapp

build:
	@go build -v -o bin/todoapp backend/main.go

db-init:
	@go run migration/main.go init
	@go run migration/main.go up

db-up:
	@go run migration/main.go up

docker-start:
	@docker-compose up -d

docker-stop:
	@docker-compose down
