package main

import (
	conversion "api-gateway/conversion-service/kitex_gen/conversion/conversionservice"
	"log"
)

func main() {
	svr := conversion.NewServer(new(ConversionServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
