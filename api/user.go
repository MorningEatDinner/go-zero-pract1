package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"

	"github.com/xiaorui/go-zero-pract1/api/internal/config"
	"github.com/xiaorui/go-zero-pract1/api/internal/handler"
	"github.com/xiaorui/go-zero-pract1/api/internal/svc"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	log.Println("数据解析开始")
	conf.MustLoad(*configFile, &c)
	log.Println("数据解析结束")

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	log.Println("创建context")
	ctx := svc.NewServiceContext(c)
	log.Println("结束创建context")
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
