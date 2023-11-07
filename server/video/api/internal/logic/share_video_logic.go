package logic

import (
	"context"
	"douniu/common/consts"
	"fmt"

	"douniu/server/video/api/internal/svc"
	"douniu/server/video/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShareVideoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShareVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShareVideoLogic {
	return &ShareVideoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShareVideoLogic) ShareVideo(req *types.ShareVideoReq) (resp *types.ShareVideoResp, err error) {
	return &types.ShareVideoResp{ShareUrl: consts.OneVideoPlayUrlPrex + fmt.Sprint(req.VideoId)}, nil

}
