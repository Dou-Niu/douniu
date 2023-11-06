package svc

import (
	"douniu/common/consts"
	"douniu/common/init_db"
	"douniu/server/favorite/rpc/favoriterpc"
	"douniu/server/relation/rpc/relationrpc"
	"douniu/server/user/model"
	"douniu/server/user/rpc/internal/config"
	"douniu/server/video/rpc/videorpc"
	"github.com/bwmarrin/snowflake"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	Snowflake   *snowflake.Node
	SqlConn     sqlx.SqlConn
	RedisClient *redis.Client
	UserModel   model.UserModel
	FavoriteRpc favoriterpc.FavoriteRpc
	VideoRpc    videorpc.VideoRpc
	RelationRpc relationrpc.RelationRpc
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
		FavoriteRpc: favoriterpc.NewFavoriteRpc(zrpc.MustNewClient(c.FavoriteRpcConf)),
		VideoRpc:    videorpc.NewVideoRpc(zrpc.MustNewClient(c.VideoRpcConf)),
		RelationRpc: relationrpc.NewRelationRpc(zrpc.MustNewClient(c.RelationRpcConf)),
	}
}
