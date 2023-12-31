package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	VideoRpcConf zrpc.RpcClientConf

	MySQLConf struct {
		DataSource string
	}
	RedisConf redis.RedisConf
	KafkaConf struct {
		Addrs []string
		Topic string
	}
}
