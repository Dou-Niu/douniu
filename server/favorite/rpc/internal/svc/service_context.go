package svc

import (
	"douniu/server/common/consts"
	"douniu/server/favorite/model"
	"douniu/server/favorite/rpc/internal/config"
	"douniu/server/video/rpc/videorpc"
	"github.com/bwmarrin/snowflake"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config        config.Config
	FavoriteModel model.FavoriteModel
	RedisClient   *redis.Redis
	VideoRpc      videorpc.VideoRpc
	Snowflake     *snowflake.Node
}

func NewServiceContext(c config.Config) *ServiceContext {
	snowflakeNode, _ := snowflake.NewNode(consts.FavoriteMachineId)
	return &ServiceContext{
		Config:        c,
		FavoriteModel: model.NewFavoriteModel(sqlx.NewMysql(c.MySQLConf.DataSource)),
		RedisClient:   redis.MustNewRedis(c.RedisConf),
		Snowflake:     snowflakeNode,
		VideoRpc:      videorpc.NewVideoRpc(zrpc.MustNewClient(c.VideoRpcConf)),
		//VideoRpc: mock.NewVideoRpc(),
	}
}
