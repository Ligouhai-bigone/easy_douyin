package rpc

import (
	"context"
	"time"

	"github.com/Ligouhai-bigone/easy_douyin/kitex_gen/userdemo"
	"github.com/Ligouhai-bigone/easy_douyin/kitex_gen/userdemo/userservice"
	"github.com/Ligouhai-bigone/easy_douyin/pkg/constants"
	"github.com/Ligouhai-bigone/easy_douyin/pkg/errno"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
)

var userClient userservice.Client

func initUserRpc() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := userservice.NewClient(
		constants.UserServiceName,
		client.WithMuxConnection(1),                       // mux
		client.WithRPCTimeout(3*time.Second),              // rpc timeout
		client.WithConnectTimeout(50*time.Millisecond),    // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		client.WithSuite(trace.NewDefaultClientSuite()),   // tracer
		client.WithResolver(r),                            // resolver
	)

	if err != nil {
		panic(err)
	}

	userClient = c
}

func Register(ctx context.Context, req *userdemo.RegisterRequest) (int64, string, error) {
	resp, err := userClient.Register(ctx, req)
	if err != nil {
		return 0, "", err
	}

	if resp.BaseResp.StatusCode != 0 {
		return 0, "", errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}

	return resp.UserId, resp.Token, nil

}

func CheckUser(ctx context.Context, req *userdemo.CheckUserRequest) (int64, string, error) {
	resp, err := userClient.CheckUser(ctx, req)

	if err != nil {
		panic(err)
	}

	if resp.BaseResp.StatusCode != 0 {
		panic(errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage))
	}

	return resp.UserId, resp.Token, nil

}

func GetUser(ctx context.Context, req *userdemo.GetUserRequest) (*userdemo.User, error) {
	user := new(userdemo.User)
	resp, err := userClient.GetUser(ctx, req)

	if err != nil {
		return user, err
	}

	if resp.BaseResp.StatusCode != 0 {
		return user, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}

	return resp.User, nil

}

func RelationAction(ctx context.Context, req *userdemo.RelationActionRequest) error {
	resp, err := userClient.RelationAction(ctx, req)
	if err != nil {
		return err
	}

	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrNo(resp.BaseResp.StatusCode, resp.GetBaseResp().StatusMessage)
	}

	return nil
}
