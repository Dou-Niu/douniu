package logic

import (
	"context"

	"douniu/server/video/api/internal/svc"
	"douniu/server/video/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FeedUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFeedUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FeedUserLogic {
	return &FeedUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FeedUserLogic) FeedUser(req *types.FeedUserReq) (resp *types.FeedResp, err error) {
	// todo: add your logic here and delete this line

	return
}
