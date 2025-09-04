package pipeline

import "fmt"

// process Asynchronous / no parallel
func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func main() {
	// Pipeline: gen → square
	// step 1
	c := gen(2, 3, 4, 5)

	// step 2
	out := square(c)

	for n := range out {
		fmt.Println(n)
	}
}
