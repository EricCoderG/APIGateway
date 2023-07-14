package main

import (
	api "api-gateway/reverse-service/kitex_gen/reverse/reverseservice"
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
		new(ReverseServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "ReverseService"}),
		server.WithRegistry(r), // register service on etcd registry
		server.WithServiceAddr(&net.TCPAddr{Port: 8891}),
	)

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
