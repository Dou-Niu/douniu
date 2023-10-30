package logic

import (
	"context"
	"douniu/server/common/errorx"
	"douniu/server/common/utils"
	"fmt"
	"github.com/pkg/errors"

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
	err = utils.DefaultGetValidParams(l.ctx, req)
	if err != nil {
		return nil, errors.Wrapf(errorx.NewCodeError(1, fmt.Sprintf("validate校验错误: %v", err)), "validate校验错误err :%v", err)
	}

	return
}
