package application

import (
	"codelytv-api/internal/mooc"
	"context"
)

type GetCoursesService struct {
	repository mooc.CourseRepository
}

func NewGetCoursesService(repository mooc.CourseRepository) GetCoursesService {
	return GetCoursesService{repository}
}

func (service *GetCoursesService) Get(context context.Context) ([]mooc.Course, error) {
	return service.repository.All(context)
}
