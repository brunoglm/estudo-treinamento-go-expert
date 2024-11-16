package events

import (
	"errors"
	"sync"
)

type EventDispatcher struct {
	handlers map[string][]EventHandlerInterface
}

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		handlers: make(map[string][]EventHandlerInterface),
	}
}

func (ed *EventDispatcher) Register(eventName string, handler EventHandlerInterface) error {
	if _, ok := ed.handlers[eventName]; ok {
		for _, h := range ed.handlers[eventName] {
			if h == handler {
				return errors.New("handler already registered")
			}
		}
	}
	ed.handlers[eventName] = append(ed.handlers[eventName], handler)
	return nil
}

func (ed *EventDispatcher) Dispatch(event EventInterface) {
	if _, ok := ed.handlers[event.GetName()]; ok {
		wg := &sync.WaitGroup{}
		for _, h := range ed.handlers[event.GetName()] {
			wg.Add(1)
			go h.Handle(event, wg)
		}
		wg.Wait()
	}
}

// remove event from dispatecher
func (ed *EventDispatcher) Remove(eventName string, handler EventHandlerInterface) {
	for _, h := range ed.handlers[eventName] {
		if h == handler {
			ed.handlers[eventName] = append(ed.handlers[eventName][:0], ed.handlers[eventName][1:]...)
		}
	}
}

func (ed *EventDispatcher) Has(eventName string, handler EventHandlerInterface) bool {
	for _, h := range ed.handlers[eventName] {
		if h == handler {
			return true
		}
	}
	return false
}

// remove all
func (ed *EventDispatcher) Clear() {
	ed.handlers = make(map[string][]EventHandlerInterface)
}
