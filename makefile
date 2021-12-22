# run backend service
run:
	@go build -v -o bin/todoapp backend/main.go
	@./bin/todoapp

run-exe:
	@go build -v -o bin/todoapp.exe backend/main.go
	@./bin/todoapp.exe

build:
	@go build -v -o bin/todoapp backend/main.go

build-exe:
	@go build -v -o bin/todoapp.exe backend/main.go

test:
	@go test -v ./backend/api
	@go test -v ./backend/datastore

db-init:
	@go run migration/main.go init
	@go run migration/main.go up

migrateup:
	migrate -path migration -database "postgresql://toped:toped@localhost:5432/todo-db?sslmode=disable" -verbose up

db-up:
	@go run migration/main.go up

docker-start:
	@docker-compose up -d

docker-stop:
	@docker-compose down

mock:
	@mockgen -destination backend/datastore/mock/datastore.go github.com/Xanvial/todo-app-go/backend/datastore DataStore