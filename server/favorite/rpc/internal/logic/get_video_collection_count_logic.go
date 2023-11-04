package logic

import (
	"context"
	"douniu/server/common/consts"
	"strconv"

	"douniu/server/favorite/rpc/internal/svc"
	"douniu/server/favorite/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetVideoCollectionCountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetVideoCollectionCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetVideoCollectionCountLogic {
	return &GetVideoCollectionCountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetVideoCollectionCountLogic) GetVideoCollectionCount(in *pb.GetVideoCollectionCountRequest) (resp *pb.GetVideoCollectionCountResponse, err error) {
	videoIdStr := strconv.Itoa(int(in.VideoId))
	resp = new(pb.GetVideoCollectionCountResponse)
	count, err := l.svcCtx.RedisClient.GetCtx(l.ctx, consts.VideoCollectCountPrefix+videoIdStr)
	if err != nil {
		l.Errorf("RedisClient ZcardCtx error: %v", err)
		return
	}
	countInt, _ := strconv.Atoi(count)
	resp.Count = int64(countInt)

	return
}
