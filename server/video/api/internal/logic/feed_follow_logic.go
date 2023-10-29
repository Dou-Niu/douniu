package logic

import (
	"context"

	"douniu/server/video/api/internal/svc"
	"douniu/server/video/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FeedFollowLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFeedFollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FeedFollowLogic {
	return &FeedFollowLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FeedFollowLogic) FeedFollow(req *types.FeedFollowReq) (resp *types.FeedResp, err error) {
	// todo: add your logic here and delete this line

	return
}
