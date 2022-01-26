package courses

import (
	"bytes"
	course2 "codelytv-api/internal/application/course"
	"codelytv-api/internal/mooc"
	"codelytv-api/internal/platform/storage/mysql"
	"codelytv-api/internal/platform/storage/storagemocks"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	validUUID   = "e8d43b09-c0b0-4864-85a6-90ee60dc3057"
	invalidUUID = "e8d43b09-c0b0-4864-85a6"
)

func TestFindHandler(t *testing.T) {
	courseID, err := mooc.NewCourseID(validUUID)
	require.NoError(t, err)
	course, err := mooc.NewCourse(validUUID, "Some course name", "4h")
	require.NoError(t, err)

	courseRepository := new(storagemocks.CourseRepository)
	courseRepository.On("Find", mock.Anything, courseID).Return(course, nil)
	findCourseService := course2.NewFindCourseService(courseRepository)

	gin.SetMode(gin.TestMode)
	server := gin.New()
	server.GET(fmt.Sprintf("%s/:id", CoursesPath), FindHandler(findCourseService))

	t.Run("given a valid id, returns the correct course", func(t *testing.T) {
		path := fmt.Sprintf("%s/%s", CoursesPath, validUUID)
		request, err := http.NewRequest(http.MethodGet, path, bytes.NewBuffer([]byte{}))
		require.NoError(t, err)

		recorder := httptest.NewRecorder()
		server.ServeHTTP(recorder, request)

		response := recorder.Result()

		assert.Equal(t, http.StatusOK, response.StatusCode)
		courseRepository.AssertExpectations(t)
	})

	t.Run("given a not valid id, returns an 400 status", func(t *testing.T) {
		path := fmt.Sprintf("%s/%s", CoursesPath, invalidUUID)
		request, err := http.NewRequest(http.MethodGet, path, bytes.NewBuffer([]byte{}))
		require.NoError(t, err)

		recorder := httptest.NewRecorder()
		server.ServeHTTP(recorder, request)
		response := recorder.Result()

		assert.Equal(t, http.StatusBadRequest, response.StatusCode)
		courseRepository.AssertExpectations(t)
	})

}
func TestFindHandlerCourseNotFound(t *testing.T) {
	courseID, err := mooc.NewCourseID(validUUID)
	require.NoError(t, err)

	courseRepository := new(storagemocks.CourseRepository)
	courseRepository.On("Find", mock.Anything, courseID).Return(mooc.Course{}, mysql.NotFoundError)
	findCourseService := course2.NewFindCourseService(courseRepository)

	gin.SetMode(gin.TestMode)
	server := gin.New()
	server.GET(fmt.Sprintf("%s/:id", CoursesPath), FindHandler(findCourseService))

	t.Run("given a valid id but no courses, returns a 404 status", func(t *testing.T) {
		path := fmt.Sprintf("%s/%s", CoursesPath, validUUID)
		request, err := http.NewRequest(http.MethodGet, path, bytes.NewBuffer([]byte{}))
		require.NoError(t, err)

		recorder := httptest.NewRecorder()
		server.ServeHTTP(recorder, request)

		response := recorder.Result()

		assert.Equal(t, http.StatusNotFound, response.StatusCode)
		courseRepository.AssertExpectations(t)
	})
}
