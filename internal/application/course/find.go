package course

import (
	"codelytv-api/internal/mooc"
	"context"
)

type FindService struct {
	repository mooc.CourseRepository
}

func NewFindService(repo mooc.CourseRepository) FindService {
	return FindService{
		repository: repo,
	}
}

func (finder *FindService) Find(context context.Context, id string) (mooc.Course, error) {
	courseID, err := mooc.NewCourseID(id)
	if err != nil {
		return mooc.Course{}, err
	}

	course, err := finder.repository.Find(context, courseID)
	if err != nil {
		return mooc.Course{}, err
	}

	return course, nil
}
