package courses

import (
	"bytes"
	"codelytv-api/internal/application/course"
	"codelytv-api/internal/platform/storage/storagemocks"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateHandler(t *testing.T) {
	// generate the mock instance of our repository and define the expected behaviour
	courseRepository := new(storagemocks.CourseRepository)
	courseRepository.On("Save", mock.Anything, mock.Anything).Return(nil)
	createCourseService := course.NewCreateService(courseRepository)

	// generate a gin test instance and register the endpoints to test
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.POST(CoursesPath, CreateHandler(createCourseService))

	t.Run("given an invalid request it return 400 status", func(t *testing.T) {
		body := getJsonBody(t, createCourseRequest{
			ID:   "a0b04b9c-717c-43f2-ae87-45fb6aa6c014",
			Name: "New Course",
		})

		request, recorder := getRequestAndRecorder(t, http.MethodPost, body)
		r.ServeHTTP(recorder, request)

		response := recorder.Result()
		assert.Equal(t, http.StatusBadRequest, response.StatusCode)
		assert.Equal(t, 0, len(courseRepository.Calls))
	})

	t.Run("given a valid request it returns a 201 status", func(t *testing.T) {
		body := getJsonBody(t, createCourseRequest{
			ID:       "a0b04b9c-717c-43f2-ae87-45fb6aa6c014",
			Name:     "Golang API",
			Duration: "2h 15m",
		})

		request, recorder := getRequestAndRecorder(t, http.MethodPost, body)
		r.ServeHTTP(recorder, request)

		response := recorder.Result()
		assert.Equal(t, http.StatusCreated, response.StatusCode)
		assert.Equal(t, 1, len(courseRepository.Calls))
	})
}

func getJsonBody(t *testing.T, data createCourseRequest) []byte {
	body, err := json.Marshal(data)
	require.NoError(t, err)
	return body
}

func getRequestAndRecorder(t *testing.T, method string, payload []byte) (*http.Request, *httptest.ResponseRecorder) {
	request, err := http.NewRequest(method, CoursesPath, bytes.NewBuffer(payload))
	require.NoError(t, err)
	return request, httptest.NewRecorder()
}
