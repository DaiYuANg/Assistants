package event

import "github.com/kataras/go-events"

func NewGlobalEventEmitter() events.EventEmmiter {
	e := events.New()
	return e
}

func NewGUIEventEmitter() events.EventEmmiter {
	e := events.New()
	return e
}

func NewBackendEventEmitter() events.EventEmmiter {
	return events.New()
}
