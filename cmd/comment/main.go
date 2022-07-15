package main

import (
	commentdemo "github.com/Ligouhai-bigone/easy_douyin/cmd/comment/kitex_gen/commentdemo/noteservice"
	"log"
)

func main() {
	svr := commentdemo.NewServer(new(NoteServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
