package nacos

import (
	"stuoj-api/infrastructure/client"
	"stuoj-api/interfaces/rpc/interceptors"
	"stuoj-common/pkg/utils"
	"stuoj-gateway/pkg/config"
)

func LoadConfig() error {
	var err error

	if err = NacosClient.GetConfig(&config.Conf.Gateway); err != nil {
		return err
	}

	utils.Expire = config.Conf.Gateway.Token.Expire
	utils.Secret = config.Conf.Gateway.Token.Secret
	utils.Refresh = config.Conf.Gateway.Token.Refresh

	client.Token = config.Conf.Gateway.Grpc.Token
	interceptors.Token = config.Conf.Gateway.Grpc.Token

	return nil
}
