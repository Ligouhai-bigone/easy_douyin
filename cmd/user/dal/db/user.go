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

func QueryUser(ctx context.Context, userName string) ([]*User, error) {
	res := make([]*User, 0)
	if err := DB.WithContext(ctx).Where("user_name = ?", userName).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
