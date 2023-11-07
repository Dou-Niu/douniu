package svc

import (
	"douniu/common/consts"
	"douniu/common/init_db"
	"douniu/server/comment/rpc/commentrpc"
	"douniu/server/favorite/rpc/favoriterpc"
	"douniu/server/user/rpc/userrpc"
	"douniu/server/video/model"
	"douniu/server/video/rpc/internal/config"
	"github.com/bwmarrin/snowflake"
	"github.com/olivere/elastic/v7"
	"github.com/zeromicro/go-zero/zrpc"

	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config         config.Config
	Snowflake      *snowflake.Node
	SqlConn        sqlx.SqlConn
	RedisClient    *redis.Client
	VideoModel     model.VideoModel
	KqPusherClient *kq.Pusher
	UserRpc        userrpc.UserRpc
	FavoriteRpc    favoriterpc.FavoriteRpc
	CommentRpc     commentrpc.CommentRpc
	ESClient       *elastic.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysqlConn := sqlx.NewMysql(c.MysqlConf.DataSource)
	snowflakeNode, _ := snowflake.NewNode(consts.VideoMachineId)
	return &ServiceContext{
		Config:         c,
		SqlConn:        mysqlConn,
		VideoModel:     model.NewVideoModel(mysqlConn, c.CacheRedis),
		Snowflake:      snowflakeNode,
		RedisClient:    init_db.InitRedis(c.RedisConf.Host, c.RedisConf.Password),
		KqPusherClient: kq.NewPusher(c.KqPusherConf.Brokers, c.KqPusherConf.Topic),
		UserRpc:        userrpc.NewUserRpc(zrpc.MustNewClient(c.UserRpcConf)),
		FavoriteRpc:    favoriterpc.NewFavoriteRpc(zrpc.MustNewClient(c.FavoriteRpcConf)),
		CommentRpc:     commentrpc.NewCommentRpc(zrpc.MustNewClient(c.CommentRpcConf)),
		ESClient:       init_db.GetESClient(c.ESConf.Host),
	}
}
