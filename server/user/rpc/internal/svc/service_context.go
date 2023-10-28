package svc

import (
	"douniu/server/common/consts"
	"douniu/server/common/init_db"
	"douniu/server/user/model"
	"douniu/server/user/rpc/internal/config"
	"github.com/bwmarrin/snowflake"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config      config.Config
	Snowflake   *snowflake.Node
	SqlConn     sqlx.SqlConn
	RedisClient *redis.Client
	UserModel   model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysqlConn := sqlx.NewMysql(c.MysqlConf.DataSource)
	snowflakeNode, _ := snowflake.NewNode(consts.UserMachineId)
	return &ServiceContext{
		Config:      c,
		SqlConn:     mysqlConn,
		UserModel:   model.NewUserModel(mysqlConn, c.CacheRedis),
		Snowflake:   snowflakeNode,
		RedisClient: init_db.InitRedis(c.RedisConf.Host, c.RedisConf.Password),
	}
}
