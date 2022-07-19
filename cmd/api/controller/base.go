package controller

import (
	"net/http"

	"github.com/Ligouhai-bigone/easy_douyin/pkg/errno"
	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int64       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SendResponse(c *gin.Context, err error, data interface{}) {
	Err := errno.ConvertErr(err)
	c.JSON(http.StatusOK, Response{
		Code:    Err.ErrCode,
		Message: Err.ErrMsg,
		Data:    data,
	})
}

type UserParam struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
}

type RegisterParam struct {
	UserId int64
	Token  string
}
