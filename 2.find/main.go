package main

import (
	"fmt"
	api "github.com/hashicorp/consul/api"
	"time"
)

const (
	consulAgentAddress = "127.0.0.1:8500"
)

// ConsulFindServer 从 consul 中发现服务
func ConsulFindServer() {
	// 创建连接 consul 服务配置
	client, err := api.NewClient(&api.Config{Address: consulAgentAddress})
	if err != nil {
		fmt.Println("consul client error : ", err)
		return
	}

	// 获取指定 service
	service, _, err := client.Agent().Service("337", nil)
	if err != nil {
		fmt.Println("Agent().Service fail", err)
		return
	}

	fmt.Println(service.Address)
	fmt.Println(service.Port)

	//只获取健康的 service
	serviceHealthy, _, err := client.Health().Service("service337", "", true, nil)
	if err != nil {
		fmt.Println("Health().Service fail", err)
		return
	}
	fmt.Println(serviceHealthy[0].Service.Address)
	fmt.Println(serviceHealthy[0].Service.Port)
}

func main() {
	ConsulFindServer()
	time.Sleep(time.Hour)
}
