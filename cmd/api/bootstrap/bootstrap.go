package bootstrap

import (
	"codelytv-api/internal/application/course"
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

	// instantiate repositories
	courseRepository := mysql.NewCourseRepository(db)

	// instantiate application services
	createCourseService := course.NewCreateService(courseRepository)
	findCourseService := course.NewFindService(courseRepository)
	// getCoursesService := course.NewGetCoursesService(courseRepository)

	srv := server.New(host, port, createCourseService, findCourseService)
	return srv.Run()
}
