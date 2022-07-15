package main

import (
	userdemo "github.com/Ligouhai-bigone/easy_douyin/cmd/user/kitex_gen/userdemo/userservice"
	"log"
)

func main() {
	svr := userdemo.NewServer(new(UserServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
