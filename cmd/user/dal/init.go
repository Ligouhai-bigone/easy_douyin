package dal

import (
	"github.com/Ligouhai-bigone/easy_douyin/cmd/user/dal/cache"
	"github.com/Ligouhai-bigone/easy_douyin/cmd/user/dal/db"
)

func Init() {
	db.Init()
	cache.Init()
}
