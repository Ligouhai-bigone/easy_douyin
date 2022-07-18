package main

import (
	"net/http"

	"github.com/Ligouhai-bigone/easy_douyin/cmd/api/controller"
	"github.com/Ligouhai-bigone/easy_douyin/cmd/api/middleware"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	r.Use(middleware.StatusLogger())

	douyin := r.Group("/douyin")

	user := douyin.Group("/user")
	user.POST("/login", controller.UserLogin)
	user.POST("/register", controller.UserRegister)

	if err := http.ListenAndServe(":8088", r); err != nil {
		klog.Fatal(err)
	}

}
