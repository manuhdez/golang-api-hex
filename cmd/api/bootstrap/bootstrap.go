package bootstrap

import (
	"codelytv-api/internal/application/course"
	"codelytv-api/internal/application/course/create"
	"codelytv-api/internal/platform/bus/inmemory"
	"codelytv-api/internal/platform/config"
	"codelytv-api/internal/platform/server"
	"codelytv-api/internal/platform/storage/mysql"
	"context"
)

func Run() error {
	env := config.GetEnv()

	db, err := mysql.Connect(env.Db)
	if err != nil {
		println("cannot connect to db", err.Error())
		return err
	}

	var (
		commandBus = inmemory.NewCommandBus()
	)

	// instantiate repositories
	courseRepository := mysql.NewCourseRepository(db)

	// instantiate application services
	createCourseService := create.NewCreateCourseService(courseRepository)
	findCourseService := application.NewFindCourseService(courseRepository)
	getCoursesService := application.NewGetCoursesService(courseRepository)

	// instantiate application command handlers
	createCourseHandler := create.NewCourseCommandHandler(createCourseService)
	commandBus.Register(create.CourseCommandType, createCourseHandler)

	ctx, srv := server.New(context.Background(), env.App.Host, env.App.Port, commandBus, findCourseService, getCoursesService)
	return srv.Run(ctx)
}
