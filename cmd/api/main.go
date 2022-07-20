package main

import (
	"net/http"

	"github.com/Ligouhai-bigone/easy_douyin/cmd/api/middleware"
	"github.com/Ligouhai-bigone/easy_douyin/cmd/api/rpc"

	"github.com/Ligouhai-bigone/easy_douyin/cmd/api/controller"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/gin-gonic/gin"
)

func Init() {
	rpc.InitRPC()
}

func main() {
	Init()
	r := gin.New()

	r.Use(middleware.StatusLogger())

	douyin := r.Group("/douyin")

	user := douyin.Group("/user")
	user.POST("/login", controller.UserLogin)
	user.POST("/register", controller.UserRegister)
	user.GET("/getUser", controller.GetUser)

	if err := http.ListenAndServe(":8080", r); err != nil {
		klog.Fatal(err)
	}

}
