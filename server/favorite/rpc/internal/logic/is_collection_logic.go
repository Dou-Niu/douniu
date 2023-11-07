package logic

import (
	"context"
	"douniu/common/consts"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"strconv"

	"douniu/server/favorite/rpc/internal/svc"
	"douniu/server/favorite/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type IsCollectionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIsCollectionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IsCollectionLogic {
	return &IsCollectionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *IsCollectionLogic) IsCollection(in *pb.IsCollectionRequest) (resp *pb.IsCollectionResponse, err error) {
	userIdStr := strconv.Itoa(int(in.UserId))
	videoIdStr := strconv.Itoa(int(in.VideoId))
	// 判断是否已经收藏
	isFavorited, err := l.svcCtx.RedisClient.ZscoreCtx(l.ctx, consts.UserCollectPrefix+userIdStr, videoIdStr)
	if err != nil && !errors.Is(err, redis.Nil) {
		l.Errorf("RedisClient ZscoreCtx error: %v", err)
		return
	}

	resp = new(pb.IsCollectionResponse)
	resp.IsCollection = isFavorited != 0
	err = nil

	return
}
