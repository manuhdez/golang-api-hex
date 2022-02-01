package mooc

import (
	"errors"
	"fmt"
)

type CourseName struct {
	value string
}

const (
	minCourseNameLength int = 6
	maxCourseNameLength int = 80
)

var ShortCourseNameError = errors.New(fmt.Sprintf("course name must be longer than %d characters", minCourseNameLength))
var LongCourseNameError = errors.New(fmt.Sprintf("course name can't be longer than %d characters", maxCourseNameLength))

func NewCourseName(value string) (CourseName, error) {
	if len(value) < minCourseNameLength {
		return CourseName{}, fmt.Errorf("%w: %v", ShortCourseNameError, value)
	}
	if len(value) > maxCourseNameLength {
		return CourseName{}, fmt.Errorf("%w: %v", LongCourseNameError, value)
	}
	return CourseName{value: value}, nil
}

func (cn CourseName) Value() string {
	return cn.value
}
