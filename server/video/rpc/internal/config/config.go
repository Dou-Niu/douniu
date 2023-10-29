package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	MysqlConf struct {
		DataSource string
	}

	CacheRedis cache.CacheConf
	RedisConf  struct {
		Host     string
		Password string
	}
	KqPusherConf struct {
		Brokers []string
		Topic   string
	}
}
