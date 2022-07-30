package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

func RedisSetKey(ctx context.Context, key string, value interface{}) error {
	if _, err := Redisclient.Get(ctx, key).Result(); err != redis.Nil {
		fmt.Println("Redis client set successfully...")
	}
	err := Redisclient.Set(ctx, key, value, 3*time.Minute).Err()
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

func RedisSetS(ctx context.Context, key string, value interface{}) error {
	err := Redisclient.SAdd(ctx, key, value).Err()
	if err != nil {
		panic(err)
	}

	return nil
}

func RedisDeleteS(ctx context.Context, key string, value interface{}) error {
	err := Redisclient.SRem(ctx, key, value).Err()
	if err != nil {
		panic(err)
	}

	return nil
}

func RedisGetS(ctx context.Context, key string) ([]string, error) {
	value, err := Redisclient.SMembers(ctx, key).Result()
	if err != nil {
		panic(err)
	}

	return value, nil
}

func RedisIsInS(ctx context.Context, key string, value interface{}) (bool, error) {
	isIn, err := Redisclient.SIsMember(ctx, key, value).Result()
	if err != nil {
		panic(err)
	}

	return isIn, nil
}
