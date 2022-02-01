package create

import (
	"codelytv-api/internal/mooc"
	"codelytv-api/kit/event"
	"context"
)

type CourseService struct {
	repository mooc.CourseRepository
	eventBus   event.Bus
}

func NewCreateCourseService(repository mooc.CourseRepository, bus event.Bus) CourseService {
	return CourseService{repository: repository, eventBus: bus}
}

func (service *CourseService) Create(ctx context.Context, id, name, duration string) error {
	course, err := mooc.NewCourse(id, name, duration)
	if err != nil {
		return err
	}

	err = service.repository.Save(ctx, course)
	if err != nil {
		return err
	}

	return service.eventBus.Publish(ctx, course.PullEvents())
}
