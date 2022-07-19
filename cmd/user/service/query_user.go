package service

import (
	"context"

	"github.com/Ligouhai-bigone/easy_douyin/cmd/user/dal/db"
	"github.com/Ligouhai-bigone/easy_douyin/pkg/errno"
)

type QueryUserService struct {
	ctx context.Context
}

func NewQueryUserService(ctx context.Context) *QueryUserService {
	return &QueryUserService{ctx: ctx}
}

func (s *QueryUserService) QueryUserId(userName string) (int64, error) {

	users, err := db.QueryUser(s.ctx, userName)
	if err != nil {
		return 0, err
	}
	if len(users) == 0 {
		return 0, errno.UserNotExistErr
	}
	u := users[0]
	return int64(u.ID), nil
}
