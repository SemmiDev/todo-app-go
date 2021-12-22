# Todo App Go
[![tests](https://github.com/SemmiDev/todo-app-go/actions/workflows/test.yml/badge.svg)](https://github.com/SemmiDev/todo-app-go/actions/workflows/test.yml)

- Make sure to run `go mod tidy` first to download all needed libraries
- To run this application, call `make run` on root
- To build only, call `make build` on root
- Calling one of above will generate executable file in `bin/` folder, that can be used to start the application
- After app binary is running, the web can be opened in [http://localhost:8080/](http://localhost:8080/)

# Demo
![demo](https://github.com/SemmiDev/todo-app-go/blob/main/.github/assets/demo.gif)
## Project Structure
```
│   app.env 
│   docker-compose.yaml
│   http.http
│   makefile
│
├───bin                                     # Binary output folder
├───backend
│   │   main.go                             # Main program of the app
│   │
│   ├───api                                 
│   │       main_test.go
│   │       middleware.go
│   │       server.go
│   │       todo.go
│   │       todo_test.go
│   │
│   ├───datastore                           
│   │   │   array.go
│   │   │   datastore.go
│   │   │   datastore_test.go
│   │   │   map.go
│   │   │   postgre.go
│   │   │
│   │   └───mock                            # Mock our data store
│   │           datastore.go
│   │
│   ├───entity                             
│   │       todo.go                         
│   │
│   ├───util                               
│   │       config.go
│   │       random.go
│   │
│   └───webstatic                           # Frontend Codes, taken from https://github.com/themaxsandelin/todo
│
└───migration                               # DB Migration data folder, contains up/down sql queries
```

## Database using Docker
1. Make sure docker already installed
2. Run docker compose using `make docker-start`
3. Wait until postgres downloaded
4. On first run, create the table using `make db-init`

## Postgres already installed
1. Update the config value in `model/config.go` to installed postgres configuration
2. create the table using `make db-init`
