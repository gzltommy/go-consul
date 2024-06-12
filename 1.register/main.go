package main

import (
	"fmt"
	consulapi "github.com/hashicorp/consul/api"
	"net/http"
)

const (
	consulAddress = "192.168.24.147:8500"
	localIp       = "192.168.24.117"
	localPort     = 81
)

func consulRegister() {
	// 获得一个 consul 客户端
	client, err := consulapi.NewClient(&consulapi.Config{Address: consulAddress})
	if err != nil {
		fmt.Println("consul client error : ", err)
		return
	}

	// 创建注册到 consul 的服务
	service := &consulapi.AgentServiceRegistration{
		ID:      "337",
		Name:    "service337",
		Tags:    []string{"testService"},
		Port:    localPort,
		Address: localIp,
		Check: &consulapi.AgentServiceCheck{
			Interval:                       "5s",
			Timeout:                        "5s",
			HTTP:                           fmt.Sprintf("http://%s:%d/", localIp, localPort),
			DeregisterCriticalServiceAfter: "30s", // 故障检查失败 30s 后 consul自动将注册服务删除
		},
	}

	// 注册服务到 consul
	err = client.Agent().ServiceRegister(service)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("you are visiting health check api"))
}

func main() {
	consulRegister()

	//定义一个 http 接口
	http.HandleFunc("/", Handler)
	err := http.ListenAndServe(fmt.Sprintf(":%d", localPort), nil)
	if err != nil {
		fmt.Println("error: ", err.Error())
	}
}
