package main

import (
	"log"

	commentdemo "github.com/Ligouhai-bigone/easy_douyin/kitex_gen/commentdemo/noteservice"
)

func main() {
	svr := commentdemo.NewServer(new(NoteServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
