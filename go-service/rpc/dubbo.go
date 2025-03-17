package rpc

import (
	"dubbo.apache.org/dubbo-go/v3"
	"dubbo.apache.org/dubbo-go/v3/client"
	"dubbo.apache.org/dubbo-go/v3/common/constant"
	"dubbo.apache.org/dubbo-go/v3/registry"
	"sync"
)

/**
 * @Author: 51130
 * @Date: 2025-03-14 15:31
 * @Description: dubbo.go - 描述该文件的功能
 * @Version: 1.0
 */

type DubboClient struct {
	cli *client.Client
}

var (
	instance *DubboClient
	once     sync.Once
	initErr  error
)

func GetDubboClient() (*DubboClient, error) {

	once.Do(func() {
		service, err := dubbo.NewInstance(
			dubbo.WithName("go-service"),
			dubbo.WithRegistry(
				registry.WithNacos(),
				registry.WithAddress("localhost:8848"),
				registry.WithRegisterInterface(),
			),
		)
		if err != nil {
			initErr = err
			panic(err)
			return
		}

		cli, err := service.NewClient(
			client.WithClientProtocolDubbo(),
			client.WithClientSerialization(constant.Hessian2Serialization),
		)

		if err != nil {
			initErr = err
			panic(err)
			return
		}

		instance = &DubboClient{
			cli: cli,
		}

	})

	if initErr != nil {
		return nil, initErr
	}

	return instance, nil
}
