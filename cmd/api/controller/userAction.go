package controller

import (
	"context"

	"github.com/Ligouhai-bigone/easy_douyin/cmd/api/rpc"
	"github.com/Ligouhai-bigone/easy_douyin/kitex_gen/userdemo"
	"github.com/Ligouhai-bigone/easy_douyin/pkg/errno"
	"github.com/gin-gonic/gin"
)

func UserAction(c *gin.Context) {
	var uc_Var UserActionParam

	err := c.ShouldBind(&uc_Var)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}

	if uc_Var.ToUserId == 0 || (uc_Var.ActionType != 1 && uc_Var.ActionType != 2) {
		SendResponse(c, errno.ParamErr, nil)
		return
	}

	err = rpc.RelationAction(context.Background(), &userdemo.RelationActionRequest{
		UserId:     uc_Var.UserId,
		ToUserId:   uc_Var.ToUserId,
		Token:      uc_Var.Token,
		ActionType: int32(uc_Var.ActionType),
	})

	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}

	SendResponse(c, errno.Success, nil)

}
