package svc

import (
	"douniu/server/comment/api/internal/config"
	"douniu/server/comment/rpc/commentrpc"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config     config.Config
	CommentRpc commentrpc.CommentRpc
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		CommentRpc: commentrpc.NewCommentRpc(zrpc.MustNewClient(c.CommentRpcConf)),
	}
}
