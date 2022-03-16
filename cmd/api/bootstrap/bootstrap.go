package bootstrap

import (
	"context"
	"time"

	application "github.com/manuhdez/golang-api-hex/internal/application/course"
	"github.com/manuhdez/golang-api-hex/internal/application/course/create"
	"github.com/manuhdez/golang-api-hex/internal/platform/bus/inmemory"
	"github.com/manuhdez/golang-api-hex/internal/platform/config"
	"github.com/manuhdez/golang-api-hex/internal/platform/server"
	"github.com/manuhdez/golang-api-hex/internal/platform/storage/mysql"
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
	courseRepository := mysql.NewCourseRepository(db, 5*time.Second)

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
