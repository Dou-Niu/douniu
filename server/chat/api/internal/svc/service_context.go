package svc

import (
	"douniu/server/chat/api/internal/config"
	"douniu/server/chat/rpc/chatrpc"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	ChatRpc chatrpc.ChatRpc
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		ChatRpc: chatrpc.NewChatRpc(zrpc.MustNewClient(c.ChatRpcConf)),
	}
}
