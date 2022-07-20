package cache

import (
	"context"
	"fmt"
)

func RedisSet(ctx context.Context, key string, value string) error {
	err := Redisclient.Set(ctx, key, value, 0).Err()
	if err != nil {
		panic(err)
	}

	fmt.Println("Redis client set successfully...")
	return nil
}

func RedisGetKey(ctx context.Context, key string) (string, error) {
	value, err := Redisclient.Get(ctx, key).Result()
	if err != nil {
		panic(err)
	}

	fmt.Println("Redis client get successfully...")
	return value, nil
}
