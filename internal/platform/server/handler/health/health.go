package health

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CheckHealthHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	}
}
