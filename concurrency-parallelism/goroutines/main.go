package goroutines

import (
	"fmt"
	"time"
)

func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(2 * time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	for i := 1; i <= 3; i++ {
		go worker(i) // async all
	}

	fmt.Println("Main done")
	time.Sleep(4 * time.Second)
}
