# Todo App Go
[![tests](https://github.com/SemmiDev/todo-app-go/actions/workflows/test.yml/badge.svg)](https://github.com/SemmiDev/todo-app-go/actions/workflows/test.yml)

- Make sure to run `go mod tidy` first to download all needed libraries
- To run this application, call `make run` on root
- To build only, call `make build` on root
- Calling one of above will generate executable file in `bin/` folder, that can be used to start the application
- After app binary is running, the web can be opened in [http://localhost:8080/](http://localhost:8080/)

# Architecture
Handler -> Data Store

# Demo
![demo](https://github.com/SemmiDev/todo-app-go/blob/main/.github/assets/demo.gif)
## Folder Structure
```
├───bin                                     # Binary output folder
├───backend
│   ├───api                                 
│   ├───datastore                           
│   │   └───mock                            # Mock our data store
│   ├───entity                            
│   ├───util                               
│   └───webstatic                           # Frontend Codes, taken from https://github.com/themaxsandelin/todo
└───migration                               # DB Migration data folder, contains up/down sql queries
```

## Database using Docker
1. Make sure docker already installed
2. Run docker compose using `make docker-start`
3. Wait until postgres downloaded
4. On first run, create the table using `make db-init`

## Postgres already installed
1. Update the config value in `util/config.go` to installed postgres configuration
2. create the table using `make db-init`
