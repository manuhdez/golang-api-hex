package mooc

import (
	"errors"
	"fmt"
)

type CourseDuration struct {
	value string
}

var EmptyCourseDurationError = errors.New("course duration field cannot be empty")

func NewCourseDuration(value string) (CourseDuration, error) {
	if value == "" {
		return CourseDuration{}, fmt.Errorf("%w: %v", EmptyCourseDurationError, value)
	}

	return CourseDuration{value: value}, nil
}

func (cd *CourseDuration) Value() string {
	return cd.value
}
