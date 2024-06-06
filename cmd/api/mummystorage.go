package main

import (
	"flag"
	"fmt"

	"mummy-storage/cmd/api/internal/handler"
	"mummy-storage/cmd/api/internal/svc"
	"mummy-storage/config"
	"mummy-storage/internal/environment"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "config.yaml", "the config file")

func main() {
	flag.Parse()
	var c *config.Config

	conf.MustLoad(*configFile, &c)
	// 设置log
	logx.MustSetup(c.LogConf)
	defer logx.Close()

	// 拿到资源管理器
	env := environment.NewServiceContext(c)

	server := rest.MustNewServer(c.API)
	defer server.Stop()
	ctx := svc.NewServiceContext(env)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.API.Host, c.API.Port)
	server.Start()
}
