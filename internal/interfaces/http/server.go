package http

import (
	"github.com/gin-gonic/gin"
	"stuoj-gateway/internal/interfaces/http/validator"
	"stuoj-gateway/pkg/config"
)

var (
	ginServer *gin.Engine
)

func InitServer() error {
	config := config.Conf.Gateway.Server

	// 创建gin实例
	ginServer = gin.Default()

	validator.SetupValidator()

	// 初始化路由
	err := InitRoute()
	if err != nil {
		return err
	}

	// 启动服务
	err = ginServer.Run(":" + config.Port)
	if err != nil {
		return err
	}

	return nil
}
