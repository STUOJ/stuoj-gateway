package app

import "stuoj-gateway/cmd/bootstrap"

func Main() {
	bootstrap.InitConfig()
	bootstrap.InitNacos()
	bootstrap.InitServer()
}
