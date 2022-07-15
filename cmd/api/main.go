package main

import (
	"github.com/Ligouhai-bigone/easy_douyin/cmd/api/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	r.Use(middleware.StatusLogger())

}
