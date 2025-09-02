package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type Redis interface {
	Set(ctx context.Context, key string, value string, expiration time.Duration) error
	Get(ctx context.Context, key string) (string, error)
	Del(ctx context.Context, key string) error
}

// Struct implement
type redisClient struct {
	client *redis.Client
}

// Constructor
func NewRedisClient(addr string, password string, db int) Redis {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	return &redisClient{client: rdb}
}

// Implement Set
func (r *redisClient) Set(ctx context.Context, key string, value string, expiration time.Duration) error {
	if expiration == 0 {
		expiration = -1
	}
	return r.client.Set(ctx, key, value, expiration).Err()
}

// Implement Get
func (r *redisClient) Get(ctx context.Context, key string) (string, error) {
	return r.client.Get(ctx, key).Result()
}

// Implement Del
func (r *redisClient) Del(ctx context.Context, key string) error {
	return r.client.Del(ctx, key).Err()
}

// Implement Close
func (r *redisClient) Close() error {
	return r.client.Close()
}
