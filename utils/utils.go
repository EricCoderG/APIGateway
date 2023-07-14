package utils

import (
	"context"
	"encoding/json"
	kClient "github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/genericclient"
	"github.com/cloudwego/kitex/pkg/generic"
	"github.com/cloudwego/kitex/pkg/loadbalance"
	etcd "github.com/kitex-contrib/registry-etcd"
	"os"
)

func GenerateClient(serviceName string) (genericclient.Client, error) {

	var err error

	// 初始化负载均衡
	lb := loadbalance.NewWeightedBalancer()

	// 初始化etcd解析器
	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	if err != nil {
		panic(err)
	}

	// 导入idl（泛化调用）
	workDir, _ := os.Getwd()
	p, err := generic.NewThriftFileProvider(workDir + "/idl/gateway.thrift")
	if err != nil {
		panic(err)
	}

	// 转换为 Thrift 泛化形式
	g, err := generic.JSONThriftGeneric(p)
	if err != nil {
		panic(err)
	}

	// 创建新的泛化客户端
	client, err := genericclient.NewClient(
		serviceName,
		g,
		kClient.WithResolver(r),
		kClient.WithLoadBalancer(lb),
	)
	if err != nil {
		panic(err)
	}

	return client, nil
}

func MakeRpcRequest(ctx context.Context, kClient genericclient.Client, methodName string, request interface{}, response interface{}) error {
	stringedReq, err := jsonStringify(request)
	if err != nil {
		panic(err)
	}

	respRpc, err := kClient.GenericCall(ctx, methodName, stringedReq)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal([]byte(respRpc.(string)), response)
	if err != nil {
		return err
	}

	return nil
}

func jsonStringify(item any) (string, error) {
	// 将 request struct 转换为 JSON 格式（这样就可以转换为 json 字符串）
	jsonForm, err := json.Marshal(&item)
	if err != nil {
		panic(err)
	}

	return string(jsonForm), nil
}
