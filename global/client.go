package global

import (
	"api-gateway/utils"
	"github.com/cloudwego/kitex/client/genericclient"
	"sync"
)

// ClientPool 存储各种客户端
type ClientPool struct {
	pool sync.Map
}

// NewClientPool 创建新的客户端池
func NewClientPool() *ClientPool {
	return &ClientPool{}
}

// GetClient 获取指定名称的客户端。如果不存在，就创建一个。
func (cp *ClientPool) GetClient(name string) (genericclient.Client, error) {
	client, ok := cp.pool.Load(name)
	if ok {
		return client.(genericclient.Client), nil
	}

	newClient, err := utils.GenerateClient(name)
	if err != nil {
		return nil, err
	}
	cp.pool.Store(name, newClient)
	return newClient, nil
}

var clientPool = NewClientPool()
