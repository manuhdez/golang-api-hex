package courses

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	application "github.com/manuhdez/golang-api-hex/internal/application/course"
	"github.com/manuhdez/golang-api-hex/internal/mooc"
	"github.com/manuhdez/golang-api-hex/internal/platform/storage/mysql"
)

func FindHandler(finder application.FindCourseService) gin.HandlerFunc {
	return func(context *gin.Context) {
		id := context.Param("id")

		response, err := finder.Find(context, id)
		if err != nil {
			switch {
			case errors.Is(err, mooc.InvalidUUIDError):
				context.JSON(http.StatusBadRequest, err.Error())
				return
			case errors.Is(err, mysql.NotFoundError):
				context.JSON(http.StatusNotFound, err.Error())
				return
			}
			context.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		context.JSON(http.StatusOK, response.ToPrimitives())
	}
}
