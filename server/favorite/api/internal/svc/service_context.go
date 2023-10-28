package svc

import (
	"douniu/server/favorite/api/internal/config"
	"douniu/server/favorite/rpc/favoriterpc"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	FavoriteRpc favoriterpc.FavoriteRpc
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		FavoriteRpc: favoriterpc.NewFavoriteRpc(zrpc.MustNewClient(c.FavoriteRpcConf)),
	}
}
