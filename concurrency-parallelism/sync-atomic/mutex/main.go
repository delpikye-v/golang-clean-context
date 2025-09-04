package mutex

// Mutex
import (
	"fmt"
	"sync"
)

type Counter struct {
	mu    sync.RWMutex // multiple read / single write
	value int
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) Get() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.value
}

func main() {
	c := &Counter{}
	var wg sync.WaitGroup

	// 5 goroutine incre counter
	for range 5 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Inc()
		}()
	}
	wg.Wait()

	fmt.Println("Final:", c.Get())
}
