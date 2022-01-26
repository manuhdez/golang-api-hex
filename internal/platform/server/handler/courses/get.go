package courses

import (
	"codelytv-api/internal/application/course"
	"codelytv-api/internal/mooc"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetHandler(service application.GetCoursesService) gin.HandlerFunc {
	return func(context *gin.Context) {
		courses, err := service.Get(context)
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
