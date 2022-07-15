package main

import (
	"context"

	"github.com/Ligouhai-bigone/easy_douyin/kitex_gen/userdemo"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// Rigister implements the UserServiceImpl interface.
func (s *UserServiceImpl) Rigister(ctx context.Context, req *userdemo.RigisterRequest) (resp *userdemo.RigisterResponse, err error) {
	// TODO: Your code here...
	return
}

// GetUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUser(ctx context.Context, req *userdemo.GetUserRequest) (resp *userdemo.GetUserResponse, err error) {
	// TODO: Your code here...
	return
}

// CheckUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CheckUser(ctx context.Context, req *userdemo.CheckUserRequest) (resp *userdemo.CheckUserResponse, err error) {
	// TODO: Your code here...
	return
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
