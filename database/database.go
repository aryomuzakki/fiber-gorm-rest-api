package database

import (
	"fmt"

	"gorm.io/gorm"
)

// refer to https://github.com/go-sql-driver/mysql#dsn-data-source-name for details about dsn
const (
	dbUser  = "root"
	dbPass  = ""
	dbHost  = "" // example = localhost:8080 or 127.0.0.1:3000
	dbName  = "db_fibergorm"
	dbParam = "charset=utf8mb4" // see more about parameter at https://github.com/go-sql-driver/mysql#parameters
)

var (
	Dsn string = fmt.Sprintf("%s:%s@tcp(%s)/%s?%s", dbUser, dbPass, dbHost, dbName, dbParam)
	DB  *gorm.DB
)
