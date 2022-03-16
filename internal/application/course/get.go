package application

import (
	"context"

	"github.com/manuhdez/golang-api-hex/internal/mooc"
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
