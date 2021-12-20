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
│   .gitignore
│   app.env
│   docker-compose.yaml
│   go.mod
│   go.sum
│   LICENSE
│   makefile
│   README.md
│   req.http
│
├───.github
│   └───workflows
│           test.yml
│
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
│   │   │   datastore_test.go
│   │   │   map.go
│   │   │   postgre.go
│   │   │   store.go
│   │   │
│   │   └───mock
│   │           datastore.go
│   │
│   ├───util
│   │       config.go
│   │       random.go
│   │
│   └───webstatic                           # Frontend Codes, taken from https://github.com/themaxsandelin/todo
│       │   favicon.ico
│       │   index.html
│       │
│       └───resources
│           ├───css
│           │       reset.min.css
│           │       style.css
│           │
│           └───js
│                   jquery-3.6.0.min.js
│                   main.js
│
├───bin                                     # Binary output folder
│       todoapp
|       todoapp.exe
│
├───migration                               # DB Migration data folder, contains up/down sql queries
│       1_create_main_table.down.sql
│       1_create_main_table.up.sql
│       main.go                             # Script to trigger sql queries
│
└───model                                  
        todo.go                             # Data model of main object
```

## Database using Docker
1. Make sure docker already installed
2. Run docker compose using `make docker-start`
3. Wait until postgres downloaded
4. On first run, create the table using `make db-init`

## Postgres already installed
1. Update the config value in `model/config.go` to installed postgres configuration
2. create the table using `make db-init`
