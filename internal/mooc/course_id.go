package mooc

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
)

type CourseID struct {
	value string
}

var InvalidUUIDError = errors.New("the course id has not the valid format")

func NewCourseID(value string) (CourseID, error) {
	id, err := uuid.Parse(value)
	if err != nil {
		return CourseID{}, fmt.Errorf("%w: %v", InvalidUUIDError, value)
	}

	return CourseID{value: id.String()}, nil
}

func (cid CourseID) Value() string {
	return cid.value
}
