package svc

import (
	"douniu/server/favorite/api/internal/config"
	"douniu/server/favorite/rpc/favoriterpc"
	"douniu/server/video/rpc/videorpc"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	FavoriteRpc favoriterpc.FavoriteRpc
	VideoRpc    videorpc.VideoRpc
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		FavoriteRpc: favoriterpc.NewFavoriteRpc(zrpc.MustNewClient(c.FavoriteRpcConf)),
		VideoRpc:    videorpc.NewVideoRpc(zrpc.MustNewClient(c.VideoRpcConf)),
	}
}
