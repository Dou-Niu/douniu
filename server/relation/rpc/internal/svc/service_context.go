package svc

import (
	"douniu/common/consts"
	"douniu/server/chat/rpc/chatrpc"
	"douniu/server/relation/model"
	"douniu/server/relation/rpc/internal/config"
	"douniu/server/user/rpc/userrpc"
	"github.com/bwmarrin/snowflake"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	RedisClient *redis.Redis
	UserRpc     userrpc.UserRpc
	ChatRpc     chatrpc.ChatRpc
	FollowModel model.FollowModel
	Snowflake   *snowflake.Node
}

func NewServiceContext(c config.Config) *ServiceContext {
	snowflakeNode, _ := snowflake.NewNode(consts.RelationMachineId)
	return &ServiceContext{
		Config:      c,
		FollowModel: model.NewFollowModel(sqlx.NewMysql(c.MySQLConf.DataSource)),
		RedisClient: redis.MustNewRedis(c.RedisConf),
		UserRpc:     userrpc.NewUserRpc(zrpc.MustNewClient(c.UserRpcConf)),
		//UserRpc: mock.UserRpc{},
		ChatRpc: chatrpc.NewChatRpc(zrpc.MustNewClient(c.ChatRpcConf)),
		//ChatRpc:   mock.ChatRpc{},
		Snowflake: snowflakeNode,
	}
}
