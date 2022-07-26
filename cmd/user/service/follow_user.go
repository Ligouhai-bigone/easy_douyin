package service

import (
	"context"

	"github.com/Ligouhai-bigone/easy_douyin/cmd/user/dal/cache"
	"github.com/Ligouhai-bigone/easy_douyin/cmd/user/dal/db"
)

type FollowService struct {
	ctx context.Context
}

func NewFollowService(ctx context.Context) *FollowService {
	return &FollowService{ctx: ctx}
}

func (f *FollowService) FollowUser(userid int64, to_userid int64) error {
	users, err := db.QueryUserbyId(f.ctx, userid)

	if err != nil {
		return err
	}
	to_users, err := db.QueryUserbyId(f.ctx, to_userid)
	if err != nil {
		return err
	}

	err = db.UpdateFollow(f.ctx, users)
	if err != nil {
		return err
	}
	err = db.UpdateFollower(f.ctx, to_users)
	if err != nil {
		return err
	}

	followkey := "follow" + string(rune(userid))
	followerkey := "follower" + string(rune(to_userid))

	//扩关注列表
	go func() {
		err = cache.RedisSetZ(f.ctx, followkey, to_userid)
	}()

	if err != nil {
		return err
	}

	// err = cache.RedisSetZ(f.ctx, followkey, to_userid)
	// if err != nil {
	// 	return err
	// }
	//扩粉丝列表
	go func() {
		err = cache.RedisSetZ(f.ctx, followerkey, userid)
	}()

	if err != nil {
		return err
	}

	return nil

}
