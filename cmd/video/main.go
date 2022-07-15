package main

import (
	videodemo "github.com/Ligouhai-bigone/easy_douyin/cmd/video/kitex_gen/videodemo/videoservice"
	"log"
)

func main() {
	svr := videodemo.NewServer(new(VideoServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
