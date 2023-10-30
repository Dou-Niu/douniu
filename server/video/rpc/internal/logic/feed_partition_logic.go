package logic

import (
	"context"

	"douniu/server/video/rpc/internal/svc"
	"douniu/server/video/rpc/types/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type FeedPartitionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFeedPartitionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FeedPartitionLogic {
	return &FeedPartitionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FeedPartitionLogic) FeedPartition(in *pb.FeedPartitionReq) (*pb.FeedResp, error) {
	// todo: add your logic here and delete this line

	return &pb.FeedResp{}, nil
}
