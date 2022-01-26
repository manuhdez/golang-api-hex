package application

import (
	"codelytv-api/internal/mooc"
	"codelytv-api/internal/platform/storage/storagemocks"
	"context"
	"errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
)

var validCourseId = "571b4eca-89b4-4b75-868c-5b9c9ad79ed7"

func TestFindService(t *testing.T) {
	course, err := mooc.NewCourse(validCourseId, "solid principles", "33h")
	require.NoError(t, err)
	courseId, err := mooc.NewCourseID(validCourseId)
	require.NoError(t, err)

	repository := new(storagemocks.CourseRepository)
	repository.On("Find", mock.Anything, courseId).Return(course, nil)

	service := NewFindCourseService(repository)

	t.Run("finds a course", func(t *testing.T) {
		result, err := service.Find(context.Background(), validCourseId)
		require.NoError(t, err)
		require.Equal(t, course, result)
	})

	t.Run("returns an error if the id is invalid", func(t *testing.T) {
		_, err := service.Find(context.Background(), "invalid")
		require.Error(t, err)
		require.Equal(t, true, errors.Is(err, mooc.InvalidUUIDError))
	})

	t.Run("returns an error if no course is found", func(t *testing.T) {
	})
}
