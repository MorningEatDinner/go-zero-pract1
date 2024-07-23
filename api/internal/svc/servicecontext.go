package svc

import (
	"github.com/xiaorui/go-zero-pract1/api/internal/middleware"
	"github.com/zeromicro/go-zero/rest"
	"log"

	"github.com/zeromicro/go-zero/zrpc"

	"github.com/xiaorui/go-zero-pract1/api/internal/config"
	"github.com/xiaorui/go-zero-pract1/rpc/userclient"
)

type ServiceContext struct {
	Config            config.Config
	UserRpc           userclient.User
	LoginVerification rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	log.Println("创建客户端")
	client := zrpc.MustNewClient(c.UserRpc)
	log.Println("结束创建客户端")
	return &ServiceContext{
		Config:            c,
		UserRpc:           userclient.NewUser(client),
		LoginVerification: middleware.NewLoginVerificationMiddleware().Handle,
	}
}
