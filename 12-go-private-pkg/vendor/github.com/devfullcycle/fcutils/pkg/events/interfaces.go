package events

import (
	"sync"
	"time"
)

type EventDispatcherInterface interface {
	Register(eventName string, handler EventHandlerInterface) error
	Dispatch(EventInterface)
	Remove(eventName string, handler EventHandlerInterface)
	Has(eventName string, handler EventHandlerInterface) bool
}

type EventHandlerInterface interface {
	Handle(event EventInterface, wg *sync.WaitGroup)
}

type EventInterface interface {
	GetDateTime() time.Time
	GetPayload() interface{}
	GetName() string
	SetPayload(interface{})
}
