package service

import (
	"context"
	"crypto/md5"
	"fmt"
	"io"

	"github.com/Ligouhai-bigone/easy_douyin/cmd/user/dal/db"
	"github.com/Ligouhai-bigone/easy_douyin/kitex_gen/userdemo"
	"github.com/Ligouhai-bigone/easy_douyin/pkg/errno"
)

type RegisterService struct {
	ctx context.Context
}

func NewRegisterService(ctx context.Context) *RegisterService {
	return &RegisterService{ctx: ctx}
}

func (s *RegisterService) Register(req *userdemo.RegisterRequest) error {
	users, err := db.QueryUser(s.ctx, req.UserName)
	if err != nil {
		return err
	}
	if len(users) != 0 {
		return errno.UserAlreadyExistErr
	}

	h := md5.New()

	if _, err = io.WriteString(h, req.Password); err != nil {
		return err
	}
	passWord := fmt.Sprintf("%x", h.Sum(nil))
	return db.Register(s.ctx, []*db.User{
		{
			UserName: req.UserName,
			PassWord: passWord,
		}})

}
