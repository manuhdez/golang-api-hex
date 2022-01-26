package command

import (
	"context"
	"errors"
)

type Bus interface {
	Dispatch(context.Context, Command) error
	Register(Type, Handler)
}

//go:generate mockery --case=snake --outpkg=commandmocks --output=commandmocks --name=Bus

// Command is the interface that all commands must implement.
type Command interface {
	Type() Type
}

var (
	ErrCommandNotFound = errors.New("command not found")
)

// Type is a type of command.
type Type string

// Handler is a function that handles a given command.
type Handler interface {
	Handle(context.Context, Command) error
}
