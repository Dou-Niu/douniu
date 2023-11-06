package logic

import (
	"context"
	"douniu/common/consts"
	"strconv"

	"douniu/server/favorite/rpc/internal/svc"
	"douniu/server/favorite/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetVideoFavoriteCountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetVideoFavoriteCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetVideoFavoriteCountLogic {
	return &GetVideoFavoriteCountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetVideoFavoriteCountLogic) GetVideoFavoriteCount(in *pb.GetVideoFavoriteCountRequest) (resp *pb.GetVideoFavoriteCountResponse, err error) {
	videoIdStr := strconv.Itoa(int(in.VideoId))
	resp = new(pb.GetVideoFavoriteCountResponse)
	count, err := l.svcCtx.RedisClient.ZcardCtx(l.ctx, consts.VideoFavoritedIdPrefix+videoIdStr)
	if err != nil {
		l.Errorf("RedisClient ZcardCtx error: %v", err)
		return
	}
	resp.Count = int64(count)

	return
}
