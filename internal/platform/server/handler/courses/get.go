package courses

import (
	"net/http"

	"github.com/gin-gonic/gin"
	application "github.com/manuhdez/golang-api-hex/internal/application/course"
	"github.com/manuhdez/golang-api-hex/internal/mooc"
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
