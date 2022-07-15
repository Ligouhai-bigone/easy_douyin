package main

import (
	"log"

	videodemo "github.com/Ligouhai-bigone/easy_douyin/kitex_gen/videodemo/videoservice"
)

func main() {
	svr := videodemo.NewServer(new(VideoServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
