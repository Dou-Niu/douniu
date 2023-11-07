package svc

import (
	"douniu/common/consts"
	"douniu/server/chat/model"
	"douniu/server/chat/rpc/internal/config"
	"github.com/bwmarrin/snowflake"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config       config.Config
	MessageModel model.MessageModel
	Snowflake    *snowflake.Node
	RedisClient  *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	snowflakeNode, _ := snowflake.NewNode(consts.ChatMachineId)
	return &ServiceContext{
		Config:       c,
		MessageModel: model.NewMessageModel(sqlx.NewMysql(c.MySQLConf.DataSource)),
		Snowflake:    snowflakeNode,
		RedisClient:  redis.MustNewRedis(c.RedisConf),
	}
}
