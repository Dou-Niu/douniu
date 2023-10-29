package logic

import (
	"context"

	"douniu/server/video/api/internal/svc"
	"douniu/server/video/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FeedHomeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFeedHomeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FeedHomeLogic {
	return &FeedHomeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FeedHomeLogic) FeedHome(req *types.FeedHomeReq) (resp *types.FeedHomeResp, err error) {
	// todo: add your logic here and delete this line

	return
}
