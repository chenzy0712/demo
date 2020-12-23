// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package simple

import "github.com/google/wire"

func InitializeEvent(msg string) (Event, error) {
	wire.Build(NewEvent, NewMessage, NewGreeter)
	return Event{}, nil
}
