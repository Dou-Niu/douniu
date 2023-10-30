package logic

import (
	"context"

	"douniu/server/video/api/internal/svc"
	"douniu/server/video/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FeedHotLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFeedHotLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FeedHotLogic {
	return &FeedHotLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FeedHotLogic) FeedHot(req *types.FeedHotReq) (resp *types.FeedHotResp, err error) {
	// todo: add your logic here and delete this line

	return
}
