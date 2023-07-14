package main

import (
	api "api-gateway/substring-service/kitex_gen/substring/api/substringservice"
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
		new(SubstringServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "SubstringService"}),
		server.WithRegistry(r), // register service on etcd registry
		server.WithServiceAddr(&net.TCPAddr{Port: 8892}),
	)

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
