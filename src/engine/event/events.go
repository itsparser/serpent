package event

import (
	"math"
	"reflect"
	"sync"
)

type EventRegister struct {
	once []interface{}
	on   []interface{}
}
type EventEmitter struct {
	eventsMutex         sync.Mutex
	events              map[string]*EventRegister
	addEventHandlers    []func(name string, handler interface{})
	removeEventHandlers []func(name string, handler interface{})
}

func (e *EventEmitter) Emit(name string, payload ...interface{}) {
	e.eventsMutex.Lock()
	defer e.eventsMutex.Unlock()
	if _, ok := e.events[name]; !ok {
		return
	}

	payloadV := make([]reflect.Value, 0)

	for _, p := range payload {
		payloadV = append(payloadV, reflect.ValueOf(p))
	}

	callHandlers := func(handlers []interface{}) {
		for _, handler := range handlers {
			handlerV := reflect.ValueOf(handler)
			handlerV.Call(payloadV[:int(math.Min(float64(handlerV.Type().NumIn()), float64(len(payloadV))))])
		}
	}

	callHandlers(e.events[name].on)
	callHandlers(e.events[name].once)

	e.events[name].once = make([]interface{}, 0)
}

func (e *EventEmitter) Once(name string, handler interface{}) {
	e.addEvent(name, handler, true)
}

func (e *EventEmitter) On(name string, handler interface{}) {
	e.addEvent(name, handler, false)
}

func (e *EventEmitter) addEventHandler(handler func(name string, handler interface{})) {
	e.addEventHandlers = append(e.addEventHandlers, handler)
}

func (e *EventEmitter) removeEventHandler(handler func(name string, handler interface{})) {
	e.removeEventHandlers = append(e.removeEventHandlers, handler)
}

func (e *EventEmitter) RemoveListener(name string, handler interface{}) {
	for _, mitm := range e.removeEventHandlers {
		mitm(name, handler)
	}
	e.eventsMutex.Lock()
	defer e.eventsMutex.Unlock()
	if _, ok := e.events[name]; !ok {
		return
	}
	handlerPtr := reflect.ValueOf(handler).Pointer()

	onHandlers := []interface{}{}
	for idx := range e.events[name].on {
		eventPtr := reflect.ValueOf(e.events[name].on[idx]).Pointer()
		if eventPtr != handlerPtr {
			onHandlers = append(onHandlers, e.events[name].on[idx])
		}
	}
	e.events[name].on = onHandlers

	onceHandlers := []interface{}{}
	for idx := range e.events[name].once {
		eventPtr := reflect.ValueOf(e.events[name].once[idx]).Pointer()
		if eventPtr != handlerPtr {
			onceHandlers = append(onceHandlers, e.events[name].once[idx])
		}
	}

	e.events[name].once = onceHandlers
}

// ListenerCount will be used to monitor the event count in EventRegistry
func (e *EventEmitter) ListenerCount(name string) int {
	count := 0
	e.eventsMutex.Lock()
	for key := range e.events {
		count += len(e.events[key].on) + len(e.events[key].once)
	}
	e.eventsMutex.Unlock()
	return count
}

func (e *EventEmitter) addEvent(name string, handler interface{}, once bool) {
	for _, mitm := range e.addEventHandlers {
		mitm(name, handler)
	}
	e.eventsMutex.Lock()
	if _, ok := e.events[name]; !ok {
		e.events[name] = &EventRegister{
			on:   make([]interface{}, 0),
			once: make([]interface{}, 0),
		}
	}
	if once {
		e.events[name].once = append(e.events[name].once, handler)
	} else {
		e.events[name].on = append(e.events[name].on, handler)
	}
	e.eventsMutex.Unlock()
}

func (e *EventEmitter) initEventEmitter() {
	e.events = make(map[string]*EventRegister)
}
