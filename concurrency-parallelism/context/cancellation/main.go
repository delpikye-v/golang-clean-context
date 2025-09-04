package cancellation

import (
	"context"
	"fmt"
	"time"
)

// conetext: use share goroutine
// management: goroutine
func worker(ctx context.Context, id int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d stopped: %v\n", id, ctx.Err())
			return
		default:
			fmt.Printf("Worker %d working...\n", id)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	// 2 workers run parallel
	go worker(ctx, 1)
	go worker(ctx, 3)

	time.Sleep(2 * time.Second)
	fmt.Println("Cancel all workers!")
	cancel() // propagate cancel all

	time.Sleep(1 * time.Second)
}
