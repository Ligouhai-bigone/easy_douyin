package db

import (
	"context"

	"github.com/Ligouhai-bigone/easy_douyin/pkg/constants"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName      string
	PassWord      string
	FollowCount   int64
	FollowerCount int64
	IsFollow      bool
}

func (u *User) TableName() string {
	return constants.UserTableName
}

func Register(ctx context.Context, users []*User) error {
	return DB.WithContext(ctx).Create(users).Error
}

func QueryUserbyName(ctx context.Context, userName string) ([]*User, error) {
	res := make([]*User, 0)
	if err := DB.WithContext(ctx).Where("user_name = ?", userName).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func QueryUserbyId(ctx context.Context, userId int64) ([]*User, error) {
	res := make([]*User, 0)
	if err := DB.WithContext(ctx).Where("ID = ?", userId).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func UpdateFollow(ctx context.Context, users []*User) error {
	user := users[0]
	err := DB.WithContext(ctx).Model(&user).Select("*").Updates(map[string]interface{}{
		"follow_count":   user.FollowCount + 1,
		"follower_count": user.FollowerCount}).Error

	if err != nil {
		return err
	}

	return nil

}

func UpdateFollower(ctx context.Context, users []*User) error {
	user := users[0]
	err := DB.WithContext(ctx).Model(&user).Select("*").Updates(map[string]interface{}{
		"follow_count":   user.FollowCount,
		"follower_count": user.FollowerCount + 1}).Error

	if err != nil {
		return err
	}

	return nil
}
