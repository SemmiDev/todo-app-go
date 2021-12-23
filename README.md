# Todo App Go
[![tests](https://github.com/SemmiDev/todo-app-go/actions/workflows/test.yml/badge.svg)](https://github.com/SemmiDev/todo-app-go/actions/workflows/test.yml)

## How to use
- Make sure to run `go mod tidy` first to download all needed libraries
- To run this application, call `make run` on root
- To build only, call `make build` on root
- Calling one of above will generate executable file in `bin` folder, that can be used to start the application
- After app binary is running, the web can be opened in [http://localhost:8080](http://localhost:8080/)

## Demo
![demo](https://github.com/SemmiDev/todo-app-go/blob/main/.github/assets/demo.gif)

## Architecture
`Handler` -> `Data Store`

## Folder Structure
```
├───bin               # Binary output folder
├───backend
│   ├───api           # Contains servers, routers, handlers, middlewares, etc                   
│   ├───datastore     # Contains several data stores that can be used                       
│   │   └───mock      # Mock our data store
│   ├───entity        # Represent table in database
│   ├───util          # Contains config for app and random utilities for testing purposes 
│   └───webstatic     # Frontend Codes, taken from https://github.com/themaxsandelin/todo
└───migration         # DB Migration data folder, contains up/down sql queries
```

## Database using Docker
1. Make sure docker already installed
2. Run docker compose using `make docker-start`
3. Wait until postgres downloaded
4. On first run, create the table using `make db-init`

## Postgres already installed
1. Update the config value in `app.env` to installed postgres configuration
2. create the table using `make db-init`

## Checklist
:heavy_check_mark: Array Store\
:heavy_check_mark: Map Store\
:heavy_check_mark: DB (Postgre SQL) Store\
:heavy_check_mark: Unit test\
:heavy_check_mark: API Test\
:heavy_check_mark: Mock DB\
:heavy_check_mark: Workflow for run tests
