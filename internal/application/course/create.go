package application

import (
	"codelytv-api/internal/mooc"
	"context"
)

type CreateCourseService struct {
	repository mooc.CourseRepository
}

func NewCreateCourseService(repository mooc.CourseRepository) CreateCourseService {
	return CreateCourseService{repository: repository}
}

func (service *CreateCourseService) Create(ctx context.Context, id, name, duration string) error {
	course, err := mooc.NewCourse(id, name, duration)
	if err != nil {
		return err
	}
	return service.repository.Save(ctx, course)
}
