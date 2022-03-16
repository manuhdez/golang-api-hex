package inmemory

import (
	"context"
	"log"

	"github.com/manuhdez/golang-api-hex/kit/command"
)

type CommandBus struct {
	handlers map[command.Type]command.Handler
}

func NewCommandBus() CommandBus {
	return CommandBus{handlers: make(map[command.Type]command.Handler)}
}

func (bus CommandBus) Register(commandType command.Type, handler command.Handler) {
	bus.handlers[commandType] = handler
}

func (bus CommandBus) Dispatch(ctx context.Context, cmd command.Command) error {
	handler, ok := bus.handlers[cmd.Type()]
	if !ok {
		return nil
	}

	hasError := make(chan error)
	go func() {
		err := handler.Handle(ctx, cmd)
		if err != nil {
			log.Printf("Error while handling %s - %s\n", cmd.Type(), err)
			hasError <- err
		}
		hasError <- nil
	}()

	return <-hasError
}
