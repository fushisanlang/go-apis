package main

import (
	"apis/app"
	"apis/app/logger"

	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	var (
		ctx = gctx.New()
	)

	logger.LogInfo("服务启动", ctx)
	app.Run()

}
