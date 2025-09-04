package synconce

import (
	"fmt"
	"sync"
)

var once sync.Once

func initConfig() {
	fmt.Println("Init config...")
}

func main() {
	for i := 0; i < 3; i++ {
		once.Do(initConfig) // init once
	}
}
