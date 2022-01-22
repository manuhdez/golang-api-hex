package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

const (
	dbUser = "codely"
	dbPass = "secret"
	dbHost = "db"
	dbPort = 3306
	dbName = "codely"
)

func Connect() (*sql.DB, error) {
	dbUri := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", dbUser, dbPass, dbHost, dbPort, dbName)
	return sql.Open("mysql", dbUri)
}
