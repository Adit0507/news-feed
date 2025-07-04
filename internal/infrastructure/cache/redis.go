package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	client *redis.Client
}

func NewRedisClient(redisURL string) *RedisClient {
	opt, err := redis.ParseURL(redisURL)
	if err != nil {
		panic(err)
	}

	client := redis.NewClient(opt)

	return &RedisClient{client: client}
}

func (r *RedisClient) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return r.client.Set(ctx, key, data, expiration).Err()
}  

func (r *RedisClient) Get(ctx context.Context, key string, dest interface {}) error {
	data, err := r.client.Get(ctx, key).Result()
	if err != nil {
		return err
	}

	return json.Unmarshal([]byte(data), dest)	
}

func (r *RedisClient) Delete(ctx context.Context, key string) error {
	return r.client.Del(ctx, key).Err()
}

func (r *RedisClient) ZAdd(ctx context.Context, key string, members ...redis.Z) error {
	return r.client.ZAdd(ctx, key, members...).Err()
}

func (r *RedisClient) ZRevRange(ctx context.Context, key string, start, stop int64) ([]string, error) {
	return r.client.ZRevRange(ctx, key, start, stop).Result()
}

func (r *RedisClient) Close() error {
	return r.client.Close()
}
