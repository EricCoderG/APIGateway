package main

import (
	length "api-gateway/length-service/kitex_gen/length/lengthservice"
	"log"
)

func main() {
	svr := length.NewServer(new(LengthServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
