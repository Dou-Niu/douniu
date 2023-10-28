package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Auth struct {
		AccessSecret       int64
		AccessTokenExpire  string
		RefreshTokenExpire string
	}
	MysqlConf struct {
		DataSource string
	}

	CacheRedis cache.CacheConf
	RedisConf  redis.RedisConf
	Credential struct {
		SecretId  string
		SecretKey string
	}
}
