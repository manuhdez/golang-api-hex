package courses

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/manuhdez/golang-api-hex/internal/application/course/create"
	"github.com/manuhdez/golang-api-hex/internal/mooc"
	"github.com/manuhdez/golang-api-hex/kit/command"
)

type createCourseRequest struct {
	ID       string `json:"id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Duration string `json:"duration" binding:"required"`
}

func CreateHandler(bus command.Bus) gin.HandlerFunc {
	return func(context *gin.Context) {
		var req createCourseRequest
		if err := context.BindJSON(&req); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := bus.Dispatch(context, create.NewCourseCommand(req.ID, req.Name, req.Duration))
		if err != nil {
			handleError(context, err)
		}

		context.Status(http.StatusCreated)
	}
}

func handleError(ctx *gin.Context, err error) {
	switch {
	case errors.Is(err, mooc.InvalidUUIDError),
		errors.Is(err, mooc.LongCourseNameError),
		errors.Is(err, mooc.ShortCourseNameError),
		errors.Is(err, mooc.EmptyCourseDurationError):

		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	default:
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}
