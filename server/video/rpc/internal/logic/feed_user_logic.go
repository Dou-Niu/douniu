package logic

import (
	"context"

	"douniu/server/video/rpc/internal/svc"
	"douniu/server/video/rpc/types/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type FeedUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFeedUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FeedUserLogic {
	return &FeedUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FeedUserLogic) FeedUser(in *pb.FeedUserReq) (*pb.FeedResp, error) {
	// todo: add your logic here and delete this line

	return &pb.FeedResp{}, nil
}
