package svc

import (
	"douniu/server/user/rpc/userrpc"
	"douniu/server/video/api/internal/config"
	"douniu/server/video/rpc/videorpc"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config   config.Config
	UserRpc  userrpc.UserRpc
	VideoRpc videorpc.VideoRpc
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		//UserRpc:  userrpc.NewUserRpc(zrpc.MustNewClient(c.UserRpcConf)),
		VideoRpc: videorpc.NewVideoRpc(zrpc.MustNewClient(c.VideoRpcConf)),
	}
}
