package fanout

import (
	"fmt"
	"time"
)

// var wg sync.WaitGroup => can use this if leak memory
// Worker function
func worker(id int, jobs <-chan string, results chan<- string) {
	for job := range jobs {
		// job = API call
		fmt.Printf("Worker %d processing API: %s\n", id, job)
		time.Sleep(1 * time.Second)
		results <- fmt.Sprintf("Worker %d done with %s", id, job)
	}
}

func main() {
	jobs := make(chan string, 5)
	results := make(chan string, 5)

	// 3 workers
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results) // limit worker CPU core * 2
	}

	// 5 API requests
	apiRequests := []string{
		"GET /users",
		"POST /login",
		"GET /orders",
		"POST /checkout",
		"PUT /profile",
	}

	// push API calls to jobs
	for _, req := range apiRequests {
		jobs <- req
	}
	// wait all workers finish
	close(jobs)

	// done
	for range apiRequests {
		fmt.Println(<-results)
	}
}
