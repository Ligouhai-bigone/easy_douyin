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

type GetUserParam struct {
	//需要在字段后写`form:xxx或者json：xxx`
	UserId int64  `form:"user_id"`
	Token  string `form:"token"`
}

type UserActionParam struct {
	//需要在字段后写`form:xxx或者json：xxx`
	UserId     int64  `json:"user_id"`
	ToUserId   int64  `json:"to_user_id"`
	Token      string `json:"token"`
	ActionType int64  `json:"action_type"`
}
