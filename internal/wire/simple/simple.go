package simple

import (
	"errors"
	"time"

	"git.kldmp.com/learning/demo/pkg/log"
)

type Message struct {
	msg string
}

func NewMessage(msg string) Message {
	return Message{msg: msg}
}

type Greeter struct {
	Message Message
	Grumpy  bool
}

func NewGreeter(M Message) Greeter {
	var grumpy bool
	if time.Now().Unix()%2 == 0 {
		grumpy = true
	}
	return Greeter{Message: M, Grumpy: grumpy}
}

func (g Greeter) Greet() Message {
	if g.Grumpy {
		return Message{"Go away!"}
	}
	return g.Message
}

type Event struct {
	Greeter Greeter
}

func NewEvent(g Greeter) (Event, error) {
	if g.Grumpy {
		return Event{}, errors.New("could not create event: event greeter is grumpy")
	}
	return Event{Greeter: g}, nil
}

func (e Event) Start() {
	msg := e.Greeter.Greet()
	log.Info("Greet:%s", msg)
}

func OldDemo() {
	message := NewMessage("hello")
	greeter := NewGreeter(message)
	event, err := NewEvent(greeter)
	if err != nil {
		return
	}
	event.Start()
}

func Demo() {
	event, err := InitializeEvent("Hello world!")

	if err == nil {
		event.Start()
	}

}
