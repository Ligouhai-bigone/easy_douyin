package service

import (
	"context"
	"strconv"

	"github.com/Ligouhai-bigone/easy_douyin/cmd/user/dal/cache"
	"github.com/Ligouhai-bigone/easy_douyin/cmd/user/dal/db"
	"github.com/Ligouhai-bigone/easy_douyin/cmd/user/pack"
	"github.com/Ligouhai-bigone/easy_douyin/kitex_gen/userdemo"
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

	users[0].FollowCount++
	to_users[0].FollowerCount++
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
		err = cache.RedisSetS(f.ctx, followkey, to_userid)
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
		err = cache.RedisSetS(f.ctx, followerkey, userid)
	}()

	if err != nil {
		return err
	}

	return nil

}

func (f *FollowService) UnFollowUser(userid int64, to_userid int64) error {
	users, err := db.QueryUserbyId(f.ctx, userid)

	if err != nil {
		return err
	}
	to_users, err := db.QueryUserbyId(f.ctx, to_userid)
	if err != nil {
		return err
	}

	users[0].FollowCount--
	to_users[0].FollowerCount--
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

	//缩关注列表
	go func() {
		err = cache.RedisDeleteS(f.ctx, followkey, to_userid)
	}()

	if err != nil {
		return err
	}

	// err = cache.RedisSetZ(f.ctx, followkey, to_userid)
	// if err != nil {
	// 	return err
	// }
	//缩粉丝列表
	go func() {
		err = cache.RedisDeleteS(f.ctx, followerkey, userid)
	}()

	if err != nil {
		return err
	}

	return nil

}

func (f *FollowService) IsFollow(userid int64, to_userid int64) bool {

	followkey := "follow" + string(rune(userid))
	isIn, err := cache.RedisIsInS(f.ctx, followkey, to_userid)
	if err != nil {
		panic(err)
	}

	return isIn

}

func (f *FollowService) IsFollower(userid int64, to_userid int64) bool {

	followkey := "follower" + string(rune(userid))
	isIn, err := cache.RedisIsInS(f.ctx, followkey, to_userid)
	if err != nil {
		panic(err)
	}

	return isIn

}

func (f *FollowService) GetFollowList(userid int64) []*userdemo.User {
	result := make([]*userdemo.User, 0)
	followkey := "follow" + string(rune(userid))
	id_list_str, err := cache.RedisGetS(f.ctx, followkey)

	if err != nil {
		panic(err)
	}

	for _, id_str := range id_list_str {
		id, err := strconv.ParseInt(id_str, 10, 64)
		if err != nil {
			panic(err)
		}
		users, err := db.QueryUserbyId(f.ctx, id)
		if err != nil {
			panic(err)
		}
		user := users[0]
		modeluser := pack.BuildUserInfoResp(id, user.UserName, user.FollowCount, user.FollowerCount, user.IsFollow)
		result = append(result, modeluser)
	}

	return result

}

func (f *FollowService) GetFollowerList(userid int64) []*userdemo.User {
	result := make([]*userdemo.User, 0)
	followkey := "follower" + string(rune(userid))
	id_list_str, err := cache.RedisGetS(f.ctx, followkey)

	if err != nil {
		panic(err)
	}

	for _, id_str := range id_list_str {
		id, err := strconv.ParseInt(id_str, 10, 64)
		if err != nil {
			panic(err)
		}
		users, err := db.QueryUserbyId(f.ctx, id)
		if err != nil {
			panic(err)
		}
		user := users[0]
		modeluser := pack.BuildUserInfoResp(id, user.UserName, user.FollowCount, user.FollowerCount, user.IsFollow)
		result = append(result, modeluser)
	}

	return result

}
