package main

import (
	"context"

	"github.com/Ligouhai-bigone/easy_douyin/kitex_gen/videodemo"
)

// VideoServiceImpl implements the last service interface defined in the IDL.
type VideoServiceImpl struct{}

// Like implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) Like(ctx context.Context, req *videodemo.LikeRequest) (resp *videodemo.LikeResponse, err error) {
	// TODO: Your code here...
	return
}

// UnLike implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) UnLike(ctx context.Context, req *videodemo.UnLikeRequest) (resp *videodemo.UnLikeResponse, err error) {
	// TODO: Your code here...
	return
}

// QueryVideoList implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) QueryVideoList(ctx context.Context, req *videodemo.QueryVideoListRequest) (resp *videodemo.QueryVideoListResponse, err error) {
	// TODO: Your code here...
	return
}

// QueryLikeVideoList implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) QueryLikeVideoList(ctx context.Context, req *videodemo.QueryLikeVideoListRequest) (resp *videodemo.QueryLikeVideoListResponse, err error) {
	// TODO: Your code here...
	return
}

// FeedVideoList implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) FeedVideoList(ctx context.Context, req *videodemo.FeedVideoListRequest) (resp *videodemo.FeedVideoListResponse, err error) {
	// TODO: Your code here...
	return
}
