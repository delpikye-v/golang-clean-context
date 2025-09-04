package synccond

import (
	"fmt"
	"sync"
	"time"
)

// wait/notify between goroutine
func main() {
	cond := sync.NewCond(&sync.Mutex{})
	ready := false

	go func() {
		cond.L.Lock()
		for !ready {
			cond.Wait() // waiting signal
		}
		fmt.Println("Worker running...")
		cond.L.Unlock()
	}()

	time.Sleep(time.Second)
	cond.L.Lock()
	ready = true
	cond.Signal() // wake up 1 goroutine
	cond.L.Unlock()

}
