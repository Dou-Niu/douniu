package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	MySQLConf struct {
		DataSource string
	}
	UserRpcConf zrpc.RpcClientConf
	RedisConf   redis.RedisConf
	ChatRpcConf zrpc.RpcClientConf
}
