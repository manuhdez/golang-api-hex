package mooc

import (
	"context"
)

type CourseRepository interface {
	Save(ctx context.Context, course Course) error
}

//go:generate mockery --case=snake --outpkg=storagemocks --output=platform/storage/storagemocks --name=CourseRepository

type Course struct {
	id       CourseID
	name     CourseName
	duration CourseDuration
}

// NewCourse creates a new course
func NewCourse(id, name, duration string) (Course, error) {
	courseID, err := NewCourseID(id)
	if err != nil {
		return Course{}, err
	}

	courseName, err := NewCourseName(name)
	if err != nil {
		return Course{}, err
	}

	courseDuration, err := NewCourseDuration(duration)
	if err != nil {
		return Course{}, err
	}

	return Course{courseID, courseName, courseDuration}, nil
}

// ID returns the course identifier
func (c *Course) ID() CourseID {
	return c.id
}

// Name returns the course name
func (c *Course) Name() CourseName {
	return c.name
}

// Duration returns the course duration
func (c *Course) Duration() CourseDuration {
	return c.duration
}
