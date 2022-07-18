package kitexmiddleware

import (
	"context"

	"github.com/cloudwego/kitex/pkg/endpoint"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
)

func CommonMiddleware(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, request, response interface{}) (err error) {
		ri := rpcinfo.GetRPCInfo(ctx)
		// get real request
		klog.Infof("real request: %+v\n", request)
		// get remote service information
		klog.Infof("remote service name: %s, remote method: %s\n", ri.To().ServiceName(), ri.To().Method())
		if err = next(ctx, request, response); err != nil {
			return err
		}
		// get real response
		klog.Infof("real response: %+v\n", response)
		return nil
	}
}
