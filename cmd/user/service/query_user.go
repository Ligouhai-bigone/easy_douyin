package service

import (
	"context"

	"github.com/Ligouhai-bigone/easy_douyin/cmd/user/dal/db"
	"github.com/Ligouhai-bigone/easy_douyin/cmd/user/pack"
	"github.com/Ligouhai-bigone/easy_douyin/kitex_gen/userdemo"
	"github.com/Ligouhai-bigone/easy_douyin/pkg/errno"
)

type QueryUserService struct {
	ctx context.Context
}

func NewQueryUserService(ctx context.Context) *QueryUserService {
	return &QueryUserService{ctx: ctx}
}

func (s *QueryUserService) QueryUserId(userName string) (int64, error) {

	users, err := db.QueryUserbyName(s.ctx, userName)
	if err != nil {
		return 0, err
	}
	if len(users) == 0 {
		return 0, errno.UserNotExistErr
	}
	u := users[0]
	return int64(u.ID), nil
}

func (s *QueryUserService) QueryUserInfo(userId int64) (*userdemo.User, error) {
	users, err := db.QueryUserbyId(s.ctx, userId)
	if err != nil {
		panic(err)
	}
	if len(users) == 0 {
		panic(errno.UserNotExistErr)
	}

	user := users[0]
	//直接将user传进去会报错
	modeluser := pack.BuildUserInfoResp(int64(user.ID), user.UserName, user.FollowCount, user.FollowerCount, user.IsFollow)
	return modeluser, nil

}
