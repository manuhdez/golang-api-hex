package server

import (
	"codelytv-api/internal/application/course"
	"codelytv-api/internal/platform/server/handler/courses"
	"codelytv-api/internal/platform/server/handler/health"
	"fmt"
	"github.com/gin-gonic/gin"
)

type Server struct {
	httpAddr string
	engine   *gin.Engine

	createCourseService course.CreateService
	findCourseService   course.FindService
}

func New(host string, port uint, createCourseService course.CreateService, findCourseService course.FindService) Server {
	srv := Server{
		engine:   gin.New(),
		httpAddr: fmt.Sprintf("%s:%d", host, port),

		createCourseService: createCourseService,
		findCourseService:   findCourseService,
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
	//s.engine.GET(courses.CoursesPath, courses.GetHandler(s.courseRepository))
	s.engine.POST(courses.CoursesPath, courses.CreateHandler(s.createCourseService))
	s.engine.GET(fmt.Sprintf("%s/:id", courses.CoursesPath), courses.FindHandler(s.findCourseService))
}
