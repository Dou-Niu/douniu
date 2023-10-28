package svc

import (
	"douniu/server/favorite/rpc/internal/config"
	"douniu/server/mock"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type ServiceContext struct {
	Config      config.Config
	RedisClient *redis.Redis
	VideoRpc    mock.VideoRpc
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		RedisClient: redis.MustNewRedis(c.RedisConf),
		//VideoRpc:    videorpc.NewVideoRpc(zrpc.MustNewClient(c.VideoRpcConf)),
		VideoRpc: mock.NewVideoRpc(),
	}
}
