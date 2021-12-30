package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

const port = ":8080"

func main() {
	server := gin.New()

	server.GET("/health", healthHandler)

	log.Fatal(server.Run(port))
}

func healthHandler(ctx *gin.Context) {
	ctx.String(http.StatusOK, "OK")
}
