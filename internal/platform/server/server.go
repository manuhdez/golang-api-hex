package server

import (
	"codelytv-api/internal/application/course"
	"codelytv-api/internal/platform/server/handler/courses"
	"codelytv-api/internal/platform/server/handler/health"
	"codelytv-api/kit/command"
	"fmt"
	"github.com/gin-gonic/gin"
)

type Server struct {
	httpAddr string
	engine   *gin.Engine

	cmdBus            command.Bus
	findCourseService application.FindCourseService
	getCoursesService application.GetCoursesService
}

func New(host string, port uint, cmdBus command.Bus, findCourseService application.FindCourseService, getCoursesService application.GetCoursesService) Server {
	srv := Server{
		engine:   gin.New(),
		httpAddr: fmt.Sprintf("%s:%d", host, port),

		cmdBus:            cmdBus,
		findCourseService: findCourseService,
		getCoursesService: getCoursesService,
	}

	srv.registerRoutes()
	return srv
}

func (s *Server) Run() error {
	println("Server running on: " + s.httpAddr)
	return s.engine.Run(s.httpAddr)
}

func (s *Server) registerRoutes() {
	s.engine.GET("/health", health.CheckHandler())
	s.engine.GET(courses.CoursesPath, courses.GetHandler(s.getCoursesService))
	s.engine.POST(courses.CoursesPath, courses.CreateHandler(s.cmdBus))
	s.engine.GET(fmt.Sprintf("%s/:id", courses.CoursesPath), courses.FindHandler(s.findCourseService))
}
