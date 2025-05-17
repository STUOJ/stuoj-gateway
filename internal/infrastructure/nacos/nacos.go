package nacos

import (
	"stuoj-common/infrastructure/nacos"
	"stuoj-gateway/pkg/config"
)

var (
	NacosClient *nacos.NacosClient
)

// Init 初始化Nacos服务
func Init() error {
	conf := &config.Conf.Nacos

	// 使用server包中的配置创建nacos客户端
	err := nacos.InitInstance(conf)
	if err != nil {
		return err
	}
	NacosClient = nacos.GetInstance()

	// 初始化nacos客户端
	if err := NacosClient.Init(); err != nil {
		return err
	}

	return nil
}
