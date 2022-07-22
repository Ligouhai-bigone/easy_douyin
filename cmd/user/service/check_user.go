package service

import (
	"context"
	"crypto/md5"
	"fmt"
	"io"

	"github.com/Ligouhai-bigone/easy_douyin/cmd/user/dal/db"
	"github.com/Ligouhai-bigone/easy_douyin/cmd/user/pack"
	"github.com/Ligouhai-bigone/easy_douyin/kitex_gen/userdemo"
	"github.com/Ligouhai-bigone/easy_douyin/pkg/errno"
)

type CheckUserService struct {
	ctx context.Context
}

func NewCheckUserService(ctx context.Context) *CheckUserService {
	return &CheckUserService{ctx: ctx}
}

func (s *CheckUserService) CheckUser(req *userdemo.CheckUserRequest) (int64, string, error) {
	h := md5.New()
	if _, err := io.WriteString(h, req.Password); err != nil {
		return 0, "", err
	}
	passWord := fmt.Sprintf("%x", h.Sum(nil))

	userName := req.UserName
	users, err := db.QueryUserbyName(s.ctx, userName)
	if err != nil {
		return 0, "", err
	}
	if len(users) == 0 {
		return 0, "", errno.UserNotExistErr
	}
	u := users[0]
	if u.PassWord != passWord {
		return 0, "", errno.LoginErr
	}

	var token string
	token = NewRedisUserService(s.ctx).RedisGetToken(int64(u.ID))
	if token == "" {
		token = pack.GenToken()
		err := NewRedisUserService(s.ctx).RedisSetToken(int64(u.ID), token)
		if err != nil {
			panic(err)
		}
	}

	return int64(u.ID), token, nil
}
