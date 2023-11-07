package config

import (
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/core/stores/cache"
)

type Config struct {
	KqConsumerConf kq.KqConf
	service.ServiceConf
	MysqlConf struct {
		DataSource string
	}

	CacheRedis cache.CacheConf
	RedisConf  struct {
		Host     string
		Password string
	}
	ESConf struct {
		Host string
	}
}
