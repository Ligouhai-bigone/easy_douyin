package controller

import (
	"context"

	"github.com/Ligouhai-bigone/easy_douyin/cmd/api/rpc"
	"github.com/Ligouhai-bigone/easy_douyin/kitex_gen/userdemo"
	"github.com/Ligouhai-bigone/easy_douyin/pkg/errno"
	"github.com/gin-gonic/gin"
)

func UserLogin(c *gin.Context) {

	var loginVar UserParam
	if err := c.ShouldBind(&loginVar); err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}

	if len(loginVar.UserName) == 0 || len(loginVar.PassWord) == 0 {
		SendResponse(c, errno.ParamErr, nil)
		return
	}

	userid, token, err := rpc.CheckUser(context.Background(), &userdemo.CheckUserRequest{
		UserName: loginVar.UserName,
		Password: loginVar.PassWord,
	})

	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}

	SendResponse(c, errno.Success, map[string]interface{}{"UserId": userid, "Token": token})
}
