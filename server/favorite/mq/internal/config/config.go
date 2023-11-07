package config

import (
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	VideoRpcConf zrpc.RpcClientConf

	MySQLConf struct {
		DataSource string
	}
	RedisConf redis.RedisConf
	KafkaConf kq.KqConf
}
