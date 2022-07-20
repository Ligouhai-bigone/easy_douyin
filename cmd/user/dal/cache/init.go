package cache

import (
	"context"

	"github.com/Ligouhai-bigone/easy_douyin/pkg/constants"
	"github.com/go-redis/redis/v8"
)

var Redisclient *redis.Client

func Init() {

	Redisclient = redis.NewClient(&redis.Options{
		Addr:     constants.Redis_URL,
		Password: constants.Redis_PassWord,
	})

	_, err := Redisclient.Ping(context.Background()).Result()
	if err != nil {
		panic("redis ping error")
	}

}
