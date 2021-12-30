package mooc

type Course struct {
	id       string
	name     string
	duration string
}

// NewCourse creates a new course
func NewCourse(id, name, duration string) Course {
	return Course{id, name, duration}
}

// ID returns the course identifier
func (c *Course) ID() string {
	return c.id
}

// Name returns the course name
func (c *Course) Name() string {
	return c.name
}

// Duration returns the course duration
func (c *Course) Duration() string {
	return c.duration
}
