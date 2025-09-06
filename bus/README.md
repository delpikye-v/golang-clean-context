# EventBus – Generic Event Emitter in Go

- `EventBus` is a lightweight, thread-safe Publish/Subscribe implementation in Go.
- It supports multiple subscribers per event and allows both normal and one-time listeners, similar to Node.js EventEmitter.

## Features

- **Subscribe**: Listen multiple times.  
- **Once**: Listen only once, automatically removed after the first trigger.  
- **UnSubscribeFn**: Remove a specific subscriber.  
- **RemoveEvent**: Remove all subscribers of an event.  
- **Notify**: Trigger events with rest parameters (`...any`).  
- Thread-safe and non-blocking (subscribers run in goroutines).

## Installation

```bash
go get github.com/delpikye-v/golang-clean-context/bus
```

## Example

```bash
package main

import (
	"fmt"
	"time"

	"github.com/delpikye-v/golang-clean-context/bus"
)

func main() {
	bus := bus.NewEventBus()

	// Subscribe multiple times
	handler := func(args ...any) {
		fmt.Println("[Always]:", args)
	}
	bus.Subscribe("greeting", handler)

	// Subscribe once
	bus.Once("greeting", func(args ...any) {
		fmt.Println("[Once]:", args)
	})

	// Trigger events
	bus.Notify("greeting", "Hi Delpi!")
	bus.Notify("greeting", "2")
	bus.Notify("greeting", "3")

	// Unsubscribe a specific handler
	bus.UnSubscribeFn("greeting", handler)
	bus.Notify("greeting", "4 (unsub)")

	// Remove entire event
	bus.RemoveEvent("greeting")
	bus.Notify("greeting", "This should not appear")

	time.Sleep(300 * time.Millisecond)
}
