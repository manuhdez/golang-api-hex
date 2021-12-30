package courses

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type createCourseRequest struct {
	ID       string `json:"id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Duration string `json:"duration" binding:"required"`
}

func CreateHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		var req createCourseRequest
		if err := context.BindJSON(&req); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		//course := mooc.NewCourse(req.ID, req.Name, req.Duration)
		context.JSON(http.StatusCreated, gin.H{"message": "Course created successfully"})
	}
}
