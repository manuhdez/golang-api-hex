package create

import (
	"context"

	"github.com/manuhdez/golang-api-hex/kit/command"
)

// CourseCommandType Holds the value of the command type.
const CourseCommandType command.Type = "command.course.create"

// CourseCommand is the command used to create a course.
// It implements the command.Command interface.
// Represents the structure of a course.
type CourseCommand struct {
	id       string
	name     string
	duration string
}

// NewCourseCommand creates a new CourseCommand.
func NewCourseCommand(id string, name string, duration string) CourseCommand {
	return CourseCommand{
		id:       id,
		name:     name,
		duration: duration,
	}
}

func (c CourseCommand) Type() command.Type {
	return CourseCommandType
}

type CourseCommandHandler struct {
	service CourseService
}

func NewCourseCommandHandler(service CourseService) CourseCommandHandler {
	return CourseCommandHandler{
		service: service,
	}
}

func (c CourseCommandHandler) Handle(ctx context.Context, cmd command.Command) error {
	courseCommand, ok := cmd.(CourseCommand)
	if !ok {
		return command.ErrCommandNotFound
	}
	return c.service.Create(ctx, courseCommand.id, courseCommand.name, courseCommand.duration)
}
