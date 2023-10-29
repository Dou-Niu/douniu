package logic

import (
	"context"

	"douniu/server/video/api/internal/svc"
	"douniu/server/video/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FeedHostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFeedHostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FeedHostLogic {
	return &FeedHostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FeedHostLogic) FeedHost(req *types.FeedHostReq) (resp *types.FeedHostResp, err error) {
	// todo: add your logic here and delete this line

	return
}
