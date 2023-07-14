package main

import (
	api "api-gateway/conversion-service/kitex_gen/conversion/conversionservice"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	"net"
)

func main() {
	// init etcd
	r, err := etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Fatal(err)
	}

	svr := api.NewServer(
		new(ConversionServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "ConversionService"}),
		server.WithRegistry(r), // register service on etcd registry
		server.WithServiceAddr(&net.TCPAddr{Port: 8890}),
	)

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
