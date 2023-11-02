package logic

import (
	"context"
	"douniu/server/common/consts"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"strconv"
	"time"

	"douniu/server/favorite/rpc/internal/svc"
	"douniu/server/favorite/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddCollectionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddCollectionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCollectionLogic {
	return &AddCollectionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddCollectionLogic) AddCollection(in *pb.AddCollectionRequest) (resp *pb.AddCollectionResponse, err error) {
	userIdStr := strconv.Itoa(int(in.UserId))
	videoIdStr := strconv.Itoa(int(in.VideoId))
	// 判断是否已经收藏
	isFavorited, err := l.svcCtx.RedisClient.ZscoreCtx(l.ctx, consts.UserCollectPrefix+userIdStr, videoIdStr)
	if err != nil && !errors.Is(err, redis.Nil) {
		l.Errorf("RedisClient ZscoreCtx error: %v", err)
		return
	}
	if isFavorited != 0 {
		return nil, errors.New("你已经收藏过了")
	}

	// 视频添加到用户的点收藏列表
	_, err = l.svcCtx.RedisClient.ZaddCtx(l.ctx, consts.UserCollectPrefix+userIdStr, time.Now().Unix(), videoIdStr)
	if err != nil {
		l.Errorf("RedisClient ZaddCtx error: %v", err)
		return
	}
	// 视频收藏数量+1
	_, err = l.svcCtx.RedisClient.IncrbyCtx(l.ctx, consts.VideoCollectCountPrefix+videoIdStr, 1)
	if err != nil {
		l.Errorf("RedisClient IncrbyCtx error: %v", err)
		return
	}

	resp = new(pb.AddCollectionResponse)
	err = nil

	return
}
