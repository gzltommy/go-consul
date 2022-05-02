package main

import (
	"fmt"
	consulapi "github.com/hashicorp/consul/api"
)

const (
	consulAgentAddress = "192.168.24.147:8500"
)

// 从 consul 中发现服务
func ConsulFindServer() {
	// 创建连接 consul 服务配置
	config := consulapi.DefaultConfig()
	config.Address = consulAgentAddress
	client, err := consulapi.NewClient(config)
	if err != nil {
		fmt.Println("consul client error : ", err)
	}

	// 获取指定 service
	service, _, err := client.Agent().Service("337", nil)
	if err == nil {
		fmt.Println(service.Address)
		fmt.Println(service.Port)
	}

	//只获取健康的 service
	//serviceHealthy, _, err := client.Health().Service("service337", "", true, nil)
	//if err == nil{
	//	fmt.Println(serviceHealthy[0].Service.Address)
	//}

}

func main() {
	ConsulFindServer()
	select {}
}
