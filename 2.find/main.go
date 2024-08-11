package main

import (
	"fmt"
	api "github.com/hashicorp/consul/api"
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

	// 获取指定 ID service
	{
		service, _, err := client.Agent().Service("c9dc568361a04350a69f864597e1e9cd", nil)
		if err != nil {
			fmt.Println("Agent().Service fail", err)
			return
		}
		fmt.Printf("server:%s:%d\n", service.Address, service.Port)
	}

	// 获取指定服务的所有健康的节点 server
	{
		serviceHealthy, _, err := client.Health().Service("hello-server", "", true, nil)
		if err != nil {
			fmt.Println("Health().Service fail", err)
			return
		}
		for _, v := range serviceHealthy {
			fmt.Printf("health server:%s:%d\n", v.Service.Address, v.Service.Port)
		}
	}

	// 获取指定服务的所有的节点 server
	{
		allServer, _, err := client.Catalog().Service("hello-server", "", nil)
		if err != nil {
			fmt.Println("Catalog().Service fail", err)
			return
		}
		for _, v := range allServer {
			fmt.Printf("server: %s:%d\n", v.ServiceAddress, v.ServicePort)
		}
	}

	// 获取 consul 上的所有 server
	{
		allServer, err := client.Agent().Services()
		if err != nil {
			fmt.Println("Agent().Services fail", err)
			return
		}
		for _, v := range allServer {
			fmt.Printf("%s: %s:%d\n", v.Service, v.Address, v.Port)
		}
	}
	// 获取 consul 上的所有 server name
	{
		allServer, _, err := client.Catalog().Services(nil)
		if err != nil {
			fmt.Println("Agent().Services fail", err)
			return
		}
		for name, tags := range allServer {
			fmt.Printf("%s:%v\n", name, tags)
		}
	}

}

func main() {
	ConsulFindServer()
}
