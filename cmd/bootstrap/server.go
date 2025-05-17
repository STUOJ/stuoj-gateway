package bootstrap

import (
	"log"
	"stuoj-gateway/internal/infrastructure/nacos"
	"stuoj-gateway/internal/interfaces/http"
)

func InitServer() {
	if err := nacos.NacosClient.Register(); err != nil {
		panic(err)
	}

	err := http.InitServer()
	if err != nil {
		log.Println("初始化服务器失败！")
		panic(err)
	}
	log.Println("初始化服务器成功")
}
