package db

import (
	"context"

	"github.com/Ligouhai-bigone/easy_douyin/pkg/constants"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserId        int64
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
