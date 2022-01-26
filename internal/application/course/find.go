package application

import (
	"codelytv-api/internal/mooc"
	"context"
)

type FindCourseService struct {
	repository mooc.CourseRepository
}

func NewFindCourseService(repo mooc.CourseRepository) FindCourseService {
	return FindCourseService{
		repository: repo,
	}
}

func (finder *FindCourseService) Find(context context.Context, id string) (mooc.Course, error) {
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
