package cache

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func RedisSet(ctx context.Context, key string, value string) error {
	if _, err := Redisclient.Get(ctx, key).Result(); err != redis.Nil {
		fmt.Println("Redis client set successfully...")
	}
	err := Redisclient.Set(ctx, key, value, 0).Err()
	if err != nil {
		panic(err)
	}

	fmt.Println("Redis client set successfully...")
	return nil
}

func RedisGetKey(ctx context.Context, key string) (string, error) {
	value, err := Redisclient.Get(ctx, key).Result()

	if err == redis.Nil {
		return value, err
	}

	if err != nil {
		panic(err)
	}

	fmt.Println("Redis client get successfully...")
	return value, nil
}
