package main

import (
	"context"
	"crypto/md5"
	"io"

	"github.com/Ligouhai-bigone/easy_douyin/cmd/user/pack"
	"github.com/Ligouhai-bigone/easy_douyin/cmd/user/service"
	"github.com/Ligouhai-bigone/easy_douyin/kitex_gen/userdemo"
	"github.com/Ligouhai-bigone/easy_douyin/pkg/errno"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// Rigister implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *userdemo.RegisterRequest) (resp *userdemo.RegisterResponse, err error) {
	// TODO: Your code here...
	resp = new(userdemo.RegisterResponse)

	if len(req.UserName) == 0 || len(req.Password) == 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	err = service.NewRegisterService(ctx).Register(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}

	userId, err := service.NewQueryUserService(ctx).QueryUserId(req.UserName)

	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResp(errno.Success)

	resp.UserId = userId
	h := md5.New()

	if _, err := io.WriteString(h, req.Password); err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}

	// passWord := fmt.Sprintf("%x", h.Sum(nil))

	resp.Token = pack.GenToken()

	//将token存入redis
	err = service.NewRedisUserService(ctx).RedisSetToken(resp.UserId, resp.Token)
	if err != nil {
		panic(err)
	}

	return resp, nil
}

// GetUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUser(ctx context.Context, req *userdemo.GetUserRequest) (resp *userdemo.GetUserResponse, err error) {
	resp = new(userdemo.GetUserResponse)
	userId := req.UserId
	token := req.Token

	if userId == 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	if token == "" {
		resp.BaseResp = pack.BuildBaseResp(errno.TokenErr)
		return resp, nil
	}

	user, err := service.NewQueryUserService(ctx).QueryUserInfo(userId)

	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResp(errno.Success)

	resp.User = user

	return resp, nil
}

// CheckUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CheckUser(ctx context.Context, req *userdemo.CheckUserRequest) (resp *userdemo.CheckUserResponse, err error) {
	resp = new(userdemo.CheckUserResponse)

	if len(req.UserName) == 0 || len(req.Password) == 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	uid, token, err := service.NewCheckUserService(ctx).CheckUser(req)

	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResp(errno.Success)

	resp.UserId = uid
	resp.Token = token

	return resp, err
}

// FollowUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) FollowUser(ctx context.Context, req *userdemo.FollowUserRequset) (resp *userdemo.FollowUserResponse, err error) {
	// TODO: Your code here...
	return
}

// UnFollowUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) UnFollowUser(ctx context.Context, req *userdemo.UnFollowUserRequset) (resp *userdemo.UnFollowUserResponse, err error) {
	// TODO: Your code here...
	return
}

// GetFollowList implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetFollowList(ctx context.Context, req *userdemo.GetFollowListRequest) (resp *userdemo.GetFollowListResponse, err error) {
	// TODO: Your code here...
	return
}

// GetFollowerList implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetFollowerList(ctx context.Context, req *userdemo.GetFollowerListRequest) (resp *userdemo.GetFollowerListResponse, err error) {
	// TODO: Your code here...
	return
}
