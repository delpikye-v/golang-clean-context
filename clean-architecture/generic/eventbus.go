package generic

// Example: GENERIC TYPE
type EventBus[T any] struct {
	subs map[string][]func(T)
}

func NewEventBus[T any]() *EventBus[T] {
	return &EventBus[T]{subs: make(map[string][]func(T))}
}

func (evt *EventBus[T]) Subscribe(eventName string, fn func(T)) {
	bus := evt.subs[eventName]
	bus = append(bus, fn)
	evt.subs[eventName] = bus
}

func (evt *EventBus[T]) Notify(eventName string, context T) {
	bus := evt.subs[eventName]
	for _, fn := range bus {
		go fn(context)
	}
}

func (evt *EventBus[T]) UnSubscribe(eventName string) {
	delete(evt.subs, eventName)
}

// func run() {
// 	bus := NewEventBus[string]()

// 	bus.Subscribe("greeting", func(msg string) {
// 		// fmt.Println("[Listener1]:", msg)
// 	})
// 	bus.Subscribe("greeting", func(msg string) {
// 		// fmt.Println("[Listener2]:", msg)
// 	})

// 	bus.Notify("greeting", "Say hi!")

// 	// time.Sleep(300 * time.Millisecond) //
// }
