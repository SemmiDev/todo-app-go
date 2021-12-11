# run backend service
run:
	@go build -v -o bin/todoapp backend/main.go
	@./bin/todoapp

build:
	@go build -v -o bin/todoapp backend/main.go

build-exe:
	@go build -v -o bin/todoapp.exe backend/main.go

run-exe:
	@go build -v -o bin/todoapp.exe backend/main.go
	@./bin/todoapp.exe

datastore-test:
	@go test -v -cover -coverprofile=cover.out -covermode=atomic ./backend/datastore

db-init:
	@go run migration/main.go init
	@go run migration/main.go up

db-up:
	@go run migration/main.go up

docker-start:
	@docker-compose up -d

docker-stop:
	@docker-compose down
