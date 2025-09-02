package redis

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	rdb := NewRedisClient("localhost:6379", "", 0)

	// Set
	err := rdb.Set(ctx, "hello", "world", 10*time.Second)
	if err != nil {
		panic(err)
	}

	// Get
	val, _ := rdb.Get(ctx, "hello")
	fmt.Println("Value:", val)

	// Del
	_ = rdb.Del(ctx, "hello")
}
