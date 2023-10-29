package logic

import (
	"context"

	"douniu/server/video/api/internal/svc"
	"douniu/server/video/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FeedPartitionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFeedPartitionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FeedPartitionLogic {
	return &FeedPartitionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FeedPartitionLogic) FeedPartition(req *types.FeedPartitionReq) (resp *types.FeedResp, err error) {
	// todo: add your logic here and delete this line

	return
}
