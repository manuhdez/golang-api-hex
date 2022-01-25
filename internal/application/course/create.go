package course

import (
	"codelytv-api/internal/mooc"
	"context"
)

type CreateService struct {
	repository mooc.CourseRepository
}

func NewCreateService(repository mooc.CourseRepository) CreateService {
	return CreateService{repository: repository}
}

func (service *CreateService) Create(ctx context.Context, id, name, duration string) error {
	course, err := mooc.NewCourse(id, name, duration)
	if err != nil {
		return err
	}
	return service.repository.Save(ctx, course)
}
