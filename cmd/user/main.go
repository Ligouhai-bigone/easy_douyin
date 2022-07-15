package main

import (
	"log"

	userdemo "github.com/Ligouhai-bigone/easy_douyin/kitex_gen/userdemo/userservice"
)

func main() {
	svr := userdemo.NewServer(new(UserServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
