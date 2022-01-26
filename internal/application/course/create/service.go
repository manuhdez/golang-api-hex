package create

import (
	"codelytv-api/internal/mooc"
	"context"
)

type CourseService struct {
	repository mooc.CourseRepository
}

func NewCreateCourseService(repository mooc.CourseRepository) CourseService {
	return CourseService{repository: repository}
}

func (service *CourseService) Create(ctx context.Context, id, name, duration string) error {
	course, err := mooc.NewCourse(id, name, duration)
	if err != nil {
		return err
	}
	return service.repository.Save(ctx, course)
}
