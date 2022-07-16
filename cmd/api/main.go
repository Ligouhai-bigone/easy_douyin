package main

import (
	"github.com/Ligouhai-bigone/easy_douyin/cmd/api/controller"
	"github.com/Ligouhai-bigone/easy_douyin/cmd/api/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	r.Use(middleware.StatusLogger())

	douyin := r.Group("/douyin")

	user := douyin.Group("/user")
	user.POST("/login", controller.UserLogin)
	user.POST("/register", controller.UserRigister)

}
