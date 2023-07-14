package main

import (
	api "api-gateway/substring-service/kitex_gen/substring/api/substringservice"
	"log"
)

func main() {
	svr := api.NewServer(new(SubstringServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
