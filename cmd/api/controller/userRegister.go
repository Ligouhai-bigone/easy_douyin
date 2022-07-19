package controller

import (
	"context"

	"github.com/Ligouhai-bigone/easy_douyin/cmd/api/rpc"

	"github.com/Ligouhai-bigone/easy_douyin/kitex_gen/userdemo"
	"github.com/Ligouhai-bigone/easy_douyin/pkg/errno"
	"github.com/gin-gonic/gin"
)

func UserRegister(c *gin.Context) {
	var registerVar UserParam
	if err := c.ShouldBind(&registerVar); err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}

	if len(registerVar.UserName) == 0 || len(registerVar.PassWord) == 0 {
		SendResponse(c, errno.ParamErr, nil)
		return
	}

	err := rpc.Register(context.Background(), &userdemo.RegisterRequest{
		UserName: registerVar.UserName,
		Password: registerVar.PassWord,
	})

	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}

	SendResponse(c, errno.Success, nil)

}
