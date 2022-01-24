package courses

import (
	"codelytv-api/internal/mooc"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetHandler(repository mooc.CourseRepository) gin.HandlerFunc {
	return func(context *gin.Context) {
		courses, err := repository.All(context)
		if err != nil {
			context.JSON(http.StatusInternalServerError, err.Error())
		}

		context.JSON(http.StatusOK, formatResponse(courses))
	}
}

func formatResponse(courses []mooc.Course) []map[string]string {
	var response []map[string]string
	for _, course := range courses {
		response = append(response, course.ToPrimitives())
	}
	return response
}
