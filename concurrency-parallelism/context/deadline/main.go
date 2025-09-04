package deadline

import (
	"context"
	"fmt"
	"time"
)

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

// share context
// ctx := context.WithValue(context.Background(), "userId")
func main() {
	deadline := time.Now().Add(2 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	go worker(ctx, 1)
	go worker(ctx, 2)

	time.Sleep(3 * time.Second) // timeout
	fmt.Println("Main finished.")
}
