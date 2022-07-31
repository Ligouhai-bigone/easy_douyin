package controller

import (
	"context"

	"github.com/Ligouhai-bigone/easy_douyin/cmd/api/rpc"
	"github.com/Ligouhai-bigone/easy_douyin/kitex_gen/userdemo"
	"github.com/Ligouhai-bigone/easy_douyin/pkg/errno"
	"github.com/gin-gonic/gin"
)

func FollowListGet(c *gin.Context) {
	var getUserParam GetUserParam
	if err := c.ShouldBind(&getUserParam); err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}

	if getUserParam.UserId == 0 || len(getUserParam.Token) == 0 {
		SendResponse(c, errno.ParamErr, nil)
		return
	}

	users, err := rpc.GetFollowList(context.Background(), &userdemo.GetFollowListRequest{
		UserId: getUserParam.UserId,
		Token:  getUserParam.Token,
	})

	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}

	SendResponse(c, errno.Success, map[string]interface{}{"Users": users})
}
