package bootstrap

import (
	"codelytv-api/internal/platform/server"
	"codelytv-api/internal/platform/storage/mysql"
)

const (
	host = "0.0.0.0"
	port = 8080
)

func Run() error {
	db, err := mysql.Connect()
	if err != nil {
		println("cannot connect to db", err.Error())
		return err
	}

	// instantiate dependencies for route handlers
	courseRepository := mysql.NewCourseRepository(db)

	srv := server.New(host, port, courseRepository)
	return srv.Run()
}
