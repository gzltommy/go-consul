package main

import (
	"fmt"
	"github.com/hashicorp/consul/api"
)

const (
	consulAgentAddress = "127.0.0.1:8500"
)

func main() {
	// 创建连接 consul 服务配置
	client, err := api.NewClient(&api.Config{Address: consulAgentAddress})
	if err != nil {
		fmt.Println("consul client error : ", err)
		return
	}
	// PUT
	{
		res, err := client.KV().Put(
			&api.KVPair{
				Key: "config/mysql",
				Value: []byte(`{
    "driver": "mysql",
    "host": "172.30.134.7",
    "port": "3306",
    "user": "root",
    "password": "1234567",
    "db_name": "forge_hero"
  }`),
			}, nil)
		if err != nil {
			fmt.Println("KV().Put.error : ", err)
			return
		}
		fmt.Println(res)
	}

	// GET
	{
		kvPair, _, err := client.KV().Get("config/mysql", nil)
		if err != nil {
			fmt.Println("KV().Get.error : ", err)
			return
		}
		fmt.Printf("key:%s,value:%s", kvPair.Key, string(kvPair.Value))
	}
}
