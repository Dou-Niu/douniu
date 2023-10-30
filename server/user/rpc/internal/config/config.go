package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf

	JWTAuth struct {
		AccessSecret       string
		AccessTokenExpire  int64
		RefreshTokenExpire int64
	}
	MysqlConf struct {
		DataSource string
	}

	CacheRedis cache.CacheConf
	RedisConf  struct {
		Host     string
		Password string
	}
	Credential struct {
		SecretId  string
		SecretKey string
	}
	Salt string
}
