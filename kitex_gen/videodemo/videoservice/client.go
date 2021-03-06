// Code generated by Kitex v0.3.2. DO NOT EDIT.

package videoservice

import (
	"context"
	"github.com/Ligouhai-bigone/easy_douyin/kitex_gen/videodemo"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	Like(ctx context.Context, req *videodemo.LikeRequest, callOptions ...callopt.Option) (r *videodemo.LikeResponse, err error)
	UnLike(ctx context.Context, req *videodemo.UnLikeRequest, callOptions ...callopt.Option) (r *videodemo.UnLikeResponse, err error)
	QueryVideoList(ctx context.Context, req *videodemo.QueryVideoListRequest, callOptions ...callopt.Option) (r *videodemo.QueryVideoListResponse, err error)
	QueryLikeVideoList(ctx context.Context, req *videodemo.QueryLikeVideoListRequest, callOptions ...callopt.Option) (r *videodemo.QueryLikeVideoListResponse, err error)
	FeedVideoList(ctx context.Context, req *videodemo.FeedVideoListRequest, callOptions ...callopt.Option) (r *videodemo.FeedVideoListResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kVideoServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kVideoServiceClient struct {
	*kClient
}

func (p *kVideoServiceClient) Like(ctx context.Context, req *videodemo.LikeRequest, callOptions ...callopt.Option) (r *videodemo.LikeResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Like(ctx, req)
}

func (p *kVideoServiceClient) UnLike(ctx context.Context, req *videodemo.UnLikeRequest, callOptions ...callopt.Option) (r *videodemo.UnLikeResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UnLike(ctx, req)
}

func (p *kVideoServiceClient) QueryVideoList(ctx context.Context, req *videodemo.QueryVideoListRequest, callOptions ...callopt.Option) (r *videodemo.QueryVideoListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.QueryVideoList(ctx, req)
}

func (p *kVideoServiceClient) QueryLikeVideoList(ctx context.Context, req *videodemo.QueryLikeVideoListRequest, callOptions ...callopt.Option) (r *videodemo.QueryLikeVideoListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.QueryLikeVideoList(ctx, req)
}

func (p *kVideoServiceClient) FeedVideoList(ctx context.Context, req *videodemo.FeedVideoListRequest, callOptions ...callopt.Option) (r *videodemo.FeedVideoListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.FeedVideoList(ctx, req)
}
