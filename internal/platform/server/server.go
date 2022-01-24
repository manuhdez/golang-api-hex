package server

import (
	"codelytv-api/internal/mooc"
	"codelytv-api/internal/platform/server/handler/courses"
	"codelytv-api/internal/platform/server/handler/health"
	"fmt"
	"github.com/gin-gonic/gin"
)

type Server struct {
	httpAddr string
	engine   *gin.Engine

	courseRepository mooc.CourseRepository
}

func New(host string, port uint, courseRepository mooc.CourseRepository) Server {
	srv := Server{
		engine:   gin.New(),
		httpAddr: fmt.Sprintf("%s:%d", host, port),

		courseRepository: courseRepository,
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
	s.engine.GET(courses.CoursesPath, courses.GetHandler(s.courseRepository))
	s.engine.POST(courses.CoursesPath, courses.CreateHandler(s.courseRepository))
	s.engine.GET(fmt.Sprintf("%s/:id", courses.CoursesPath), courses.FindHandler(s.courseRepository))
}
