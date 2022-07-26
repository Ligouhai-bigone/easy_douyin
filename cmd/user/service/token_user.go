package service

import (
	"context"

	"github.com/Ligouhai-bigone/easy_douyin/cmd/user/dal/cache"
)

type TokenUserService struct {
	ctx context.Context
}

func NewTokenUserService(ctx context.Context) *TokenUserService {
	return &TokenUserService{ctx: ctx}
}

func (s *TokenUserService) SetToken(userid int64, token string) error {

	err := cache.RedisSetKey(s.ctx, string(rune(userid)), token)
	if err != nil {
		panic(err)
	}

	return nil
}

func (s *TokenUserService) GetToken(userid int64) string {

	value, err := cache.RedisGetKey(s.ctx, string(rune(userid)))
	if err != nil {
		return ""
	}

	return value
}
