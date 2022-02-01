package server

import (
	"codelytv-api/internal/application/course"
	"codelytv-api/internal/platform/server/handler/courses"
	"codelytv-api/internal/platform/server/handler/health"
	"codelytv-api/kit/command"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type Server struct {
	httpAddr string
	engine   *gin.Engine

	cmdBus            command.Bus
	findCourseService application.FindCourseService
	getCoursesService application.GetCoursesService
}

func New(ctx context.Context, host string, port uint, cmdBus command.Bus, findCourseService application.FindCourseService, getCoursesService application.GetCoursesService) (context.Context, Server) {
	srv := Server{
		engine:   gin.Default(),
		httpAddr: fmt.Sprintf("%s:%d", host, port),

		cmdBus:            cmdBus,
		findCourseService: findCourseService,
		getCoursesService: getCoursesService,
	}

	srv.registerRoutes()
	return serverContext(ctx), srv
}

func (s *Server) Run(ctx context.Context) error {
	println("Server running on: " + s.httpAddr)

	server := &http.Server{
		Addr:    s.httpAddr,
		Handler: s.engine,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Server error: ", err)
		}
	}()

	<-ctx.Done()

	ctxShutdown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return server.Shutdown(ctxShutdown)
}

func (s *Server) registerRoutes() {
	s.engine.GET("/health", health.CheckHandler())
	s.engine.GET(courses.CoursesPath, courses.GetHandler(s.getCoursesService))
	s.engine.POST(courses.CoursesPath, courses.CreateHandler(s.cmdBus))
	s.engine.GET(fmt.Sprintf("%s/:id", courses.CoursesPath), courses.FindHandler(s.findCourseService))
}

func serverContext(ctx context.Context) context.Context {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	ctx, cancel := context.WithCancel(ctx)
	go func() {
		<-c
		cancel()
	}()
	return ctx
}
