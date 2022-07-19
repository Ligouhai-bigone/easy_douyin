package main

import (
	"net"

	"github.com/Ligouhai-bigone/easy_douyin/cmd/user/dal"
	userdemo "github.com/Ligouhai-bigone/easy_douyin/kitex_gen/userdemo/userservice"
	"github.com/Ligouhai-bigone/easy_douyin/pkg/constants"
	kitexmiddleware "github.com/Ligouhai-bigone/easy_douyin/pkg/kitex_middleware"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func Init() {
	dal.Init()
}

func main() {
	r, err := etcd.NewEtcdRegistry([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8889")
	if err != nil {
		panic(err)
	}
	Init()
	svr := userdemo.NewServer(new(UserServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constants.UserServiceName}),
		server.WithMiddleware(kitexmiddleware.CommonMiddleware),
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}), // limit
		server.WithServiceAddr(addr),                                       // address
		server.WithMuxTransport(),                                          // Multiplex
		server.WithRegistry(r),                                             // registry

	)
	err = svr.Run()
	if err != nil {
		klog.Fatal(err)
	}
}
