package main

import (
	reverse "api-gateway/reverse-service/kitex_gen/reverse/reverseservice"
	"log"
)

func main() {
	svr := reverse.NewServer(new(ReverseServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
