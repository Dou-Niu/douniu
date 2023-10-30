package logic

import (
	"context"

	"douniu/server/video/rpc/internal/svc"
	"douniu/server/video/rpc/types/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type FeedHotLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFeedHotLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FeedHotLogic {
	return &FeedHotLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FeedHotLogic) FeedHot(in *pb.FeedHotReq) (*pb.FeedHotResp, error) {
	// todo: add your logic here and delete this line

	return &pb.FeedHotResp{}, nil
}
