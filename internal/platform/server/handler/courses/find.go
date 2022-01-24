package courses

import (
	"codelytv-api/internal/mooc"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FindHandler(repository mooc.CourseRepository) gin.HandlerFunc {
	return func(context *gin.Context) {
		id := context.Param("id")
		courseID, err := mooc.NewCourseID(id)
		if err != nil {
			context.JSON(http.StatusBadRequest, err.Error())
			return
		}

		course, err := repository.Find(context, courseID)
		if err != nil {
			context.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		context.JSON(http.StatusOK, course.ToPrimitives())
	}
}
