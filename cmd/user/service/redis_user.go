package service

import (
	"context"

	"github.com/Ligouhai-bigone/easy_douyin/cmd/user/dal/cache"
)

type RedisUserService struct {
	ctx context.Context
}

func NewRedisUserService(ctx context.Context) *QueryUserService {
	return &QueryUserService{ctx: ctx}
}

func (s *QueryUserService) RedisSetToken(userid int64, token string) error {

	err := cache.RedisSet(s.ctx, string(rune(userid)), token)
	if err != nil {
		panic(err)
	}

	return nil
}

func (s *QueryUserService) RedisGetToken(userid int64) string {

	value, err := cache.RedisGetKey(s.ctx, string(rune(userid)))
	if err != nil {
		return ""
	}

	return value
}
