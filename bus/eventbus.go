package bus
// eventbus with rest params

import (
	"fmt"
	"sync"
)

type EventBus struct {
	mu   sync.Mutex
	subs map[string][]func(args ...any)
}

func NewEventBus() *EventBus {
	return &EventBus{subs: make(map[string][]func(args ...any))}
}

func (evt *EventBus) Subscribe(eventName string, fn func(args ...any)) {
	evt.mu.Lock()
	defer evt.mu.Unlock()
	evt.subs[eventName] = append(evt.subs[eventName], fn)
}

func (evt *EventBus) UnSubscribeFn(eventName string, target func(args ...any)) {
	evt.mu.Lock()
	defer evt.mu.Unlock()

	if handlers, ok := evt.subs[eventName]; ok {
		// clone new array
		newHandlers := []func(args ...any){}
		for _, h := range handlers {
			// remove
			if fmt.Sprintf("%p", h) != fmt.Sprintf("%p", target) {
				newHandlers = append(newHandlers, h)
			}
		}
		evt.subs[eventName] = newHandlers
	}
}

func (evt *EventBus) Once(eventName string, target func(args ...any)) {
	var wrapper func(args ...any)
	wrapper = func(args ...any) {
		target(args...)
		evt.UnSubscribeFn(eventName, wrapper)
	}
	evt.Subscribe(eventName, wrapper)
}

// RemoveEvent
func (evt *EventBus) RemoveEvent(eventName string) {
	evt.mu.Lock()
	defer evt.mu.Unlock()
	delete(evt.subs, eventName)
}

// Notify: notify any
func (evt *EventBus) Notify(eventName string, args ...any) {
	bus := evt.subs[eventName]
	for _, fn := range bus {
		go fn(args...)
	}
}

// func main() {
// 	bus := NewEventBus()

// 	// Subscribe
// 	handler := func(args ...any) {
// 		fmt.Println("[Always]:", args)
// 	}
// 	bus.Subscribe("greeting", handler)

// 	// Subscribe once
// 	bus.Once("greeting", func(args ...any) {
// 		fmt.Println("[Once]:", args)
// 	})

// 	// notify event
// 	bus.Notify("greeting", "Hi Delpi!")
// 	bus.Notify("greeting", "2")
// 	bus.Notify("greeting", "3")

// 	// remove
// 	bus.UnSubscribeFn("greeting", handler)
// 	bus.Notify("greeting", "4 (unsub)")

// 	bus.RemoveEvent("greeting")
// 	bus.Notify("greeting", "This should not appear")

// 	time.Sleep(300 * time.Millisecond)
// }
