# fiber-gorm-rest-api
Learn to build REST API with [GO Language](https://go.dev) using [GORM](https://gorm.io/) and [Fiber](https://gofiber.io) Framework and MYSQL database

## Endpoints

| HTTP Method | Endpoint | Description |
| ---- | ---- | ---- |
| GET | `/addresses` | Get all Addresses |
| GET | `/address/:id` | Get Address by id |
| POST | `/address` | Create new Address |
| PUT | `/address` | Update Address by id (specified in request body) |
| DELETE | `/address/:id` | Delete Address by id |

## Data 

| Field | Type |
| ---- | ---- |
| id | int |
| title | string |
| description | string |
| coordinate | string |

## Build Steps
- clone repository
- create new mysql database
- import database from *addresses.sql* file
> Note: Uncomment migrate() function in main.go to use GORM AutoMigrate (a different mysql type will be used)
- change db configuration in *database/database.go* file
- run these commands from root folder
  ```cmd
  go mod tidy
  ```
  *to add missing package(s) (if necessary)*

  ```cmd
  go run main.go
  ```
  *to start server at https://localhost:4000*