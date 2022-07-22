package pack

import "github.com/Ligouhai-bigone/easy_douyin/kitex_gen/userdemo"

func BuildUserInfoResp(userid int64, username string, followcount int64, followercount int64, isfollow bool) *userdemo.User {

	user := &userdemo.User{
		UserId:        userid,
		UserName:      username,
		FollowCount:   followcount,
		FollowerCount: followercount,
		IsFollow:      isfollow,
	}

	return user
}
