package bootstrap

import (
	"codelytv-api/internal/platform/server"
	"database/sql"
	"fmt"
)

const (
	host   = "localhost"
	port   = 8080
	dbUser = "root"
	dbPass = "root"
	dbHost = "localhost"
	dbPort = 3306
	dbName = "codely"
)

func Run() error {
	dbUri := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", dbUri)
	if err != nil {
		println("cannot connect to db", error.Error())
		return err
	}

	// instantiate dependencies for route handlers
	// courseRepository := mysql.NewCourseRepository(db)

	srv := server.New(host, port)
	return srv.Run()
}
