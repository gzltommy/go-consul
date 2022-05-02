package main

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"github.com/hashicorp/consul/api/watch"
)

// watch 包的使用方法为：
//	1）使用 watch.Parse(查询参数)生成 Plan，
//	2）绑定 Plan 的 handler，
//	3）运行 Plan

// watch 类型:
// services：所有服务
// service：单个服务
// key：单个键值
// nodes：节点
// keyprefix：指定前缀的一批键值

func main() {
	//watchKey()
	//watchServices()
	watchServer()
}

func watchKey() {
	params := map[string]interface{}{
		"type":       "key",
		"datacenter": "dc1",
		//"token":      "12345",
		"key": "config",
	}
	plan, err := watch.Parse(params)
	if err != nil {
		fmt.Printf("\n Parse err: %v", err)
	}
	defer plan.Stop()
	if plan.Datacenter != "dc1" {
		fmt.Printf("\n Datacenter error: %#v", plan)
		return
	}
	//if p.Token != "12345" {
	//	log.Fatalf("3Bad: %#v", p)
	//}
	if plan.Type != "key" {
		fmt.Printf("\n type error: %#v", plan)
		return
	}

	plan.Handler = func(idx uint64, data interface{}) {
		v := data.(*api.KVPair)
		fmt.Printf("\n 数据变化 idx:%d, [key:%s] -> v:%s", idx, v.Key, string(v.Value))
	}

	if err = plan.Run("192.168.24.147:8500"); err != nil {
		fmt.Printf("\n Run error: %v", err)
		return
	}
}

func watchServices() {
	params := map[string]interface{}{
		"type":       "services",
		"datacenter": "dc1",
	}
	plan, err := watch.Parse(params)
	if err != nil {
		fmt.Printf("\n Parse err: %v", err)
	}
	defer plan.Stop()
	if plan.Datacenter != "dc1" {
		fmt.Printf("\n Datacenter error: %#v", plan)
		return
	}

	plan.Handler = func(idx uint64, data interface{}) {
		services := data.(map[string][]string)
		for k, v := range services {
			fmt.Printf("\n %s:%+v", k, v)
		}
	}

	if err = plan.Run("192.168.24.147:8500"); err != nil {
		fmt.Printf("\n Run error: %v", err)
		return
	}
}

func watchServer() {
	params := map[string]interface{}{
		"type":       "service",
		"service":    "service337",
		"datacenter": "dc1",
	}
	plan, err := watch.Parse(params)
	if err != nil {
		fmt.Printf("\n Parse err: %v", err)
	}
	defer plan.Stop()
	if plan.Datacenter != "dc1" {
		fmt.Printf("\n Datacenter error: %#v", plan)
		return
	}

	plan.Handler = func(idx uint64, data interface{}) {
		services := data.([]*api.ServiceEntry)
		for _, v := range services {
			// 这里是单个 service 变化时需要做的逻辑，可以自己添加，或在外部写一个类似handler的函数传进来
			fmt.Printf("\n service %s 已变化", v.Service.Service)
			// 打印 service 的状态
			fmt.Println("\n service status: ", v.Checks.AggregatedStatus())
		}
	}

	if err = plan.Run("192.168.24.147:8500"); err != nil {
		fmt.Printf("\n Run error: %v", err)
		return
	}
}
