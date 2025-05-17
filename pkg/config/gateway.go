package config

import "stuoj-common/pkg/config"

type GatewayConfig struct {
	Server config.ServerConf `yaml:"server" json:"server"`
	Grpc   config.GrpcConf   `yaml:"grpc" json:"grpc"`
	Token  TokenConf         `yaml:"token" json:"token"`
}

func (c *GatewayConfig) Default() {
	c.Server.Default()
	c.Grpc.Default()
	c.Token.Default()
}
