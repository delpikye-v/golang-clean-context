package fanin

import (
	"fmt"
	"time"
)

func producer(id int, out chan<- string) {
	for i := 1; i <= 3; i++ {
		out <- fmt.Sprintf("Producer %d: item %d", id, i)
		time.Sleep(300 * time.Millisecond)
	}
}

func main() {
	out := make(chan string)

	// 2 producer
	go producer(1, out)
	go producer(2, out)

	// result from 2 (fan-in)
	count := 1
	for item := range out {
		fmt.Printf("Result %d => %s\n", count, item)
		count++
	}

	// Result 1 => Producer 1: item 1
	// Result 2 => Producer 2: item 1
	// Result 3 => Producer 1: item 2
	// Result 4 => Producer 2: item 2
	// Result 5 => Producer 1: item 3
	// Result 6 => Producer 2: item 3
}
