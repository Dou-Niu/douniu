package logic

import (
	"context"

	"douniu/server/video/rpc/internal/svc"
	"douniu/server/video/rpc/types/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type FeedHomeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFeedHomeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FeedHomeLogic {
	return &FeedHomeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FeedHomeLogic) FeedHome(in *pb.FeedHomeReq) (*pb.FeedHomeResp, error) {
	// todo: add your logic here and delete this line

	return &pb.FeedHomeResp{}, nil
}
