package atomic

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
use it for simple logic: flag, counter
*/
func main() {
	var counter int64
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt64(&counter, 1)
		}()
	}

	wg.Wait()
	fmt.Println("Final:", counter)
}
