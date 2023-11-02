package logic

import (
	"context"
	"douniu/server/common/consts"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"strconv"

	"douniu/server/favorite/rpc/internal/svc"
	"douniu/server/favorite/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelCollectionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelCollectionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelCollectionLogic {
	return &DelCollectionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelCollectionLogic) DelCollection(in *pb.DelCollectionRequest) (resp *pb.DelCollectionResponse, err error) {
	userIdStr := strconv.Itoa(int(in.UserId))
	videoIdStr := strconv.Itoa(int(in.VideoId))
	// 判断是否已经收藏
	isFavorited, err := l.svcCtx.RedisClient.ZscoreCtx(l.ctx, consts.UserCollectPrefix+userIdStr, videoIdStr)
	if err != nil && !errors.Is(err, redis.Nil) {
		l.Errorf("RedisClient ZscoreCtx error: %v", err)
		return
	}
	if isFavorited == 0 {
		return nil, errors.New("你还没有收藏过")
	}

	// 删除收藏
	if _, err = l.svcCtx.RedisClient.ZremCtx(l.ctx, consts.UserCollectPrefix+userIdStr, videoIdStr); err != nil {
		l.Errorf("RedisClient ZremCtx error: %v", err)
		return
	}

	// 视频收藏数-1
	if _, err = l.svcCtx.RedisClient.DecrCtx(l.ctx, consts.VideoCollectCountPrefix+videoIdStr); err != nil {
		l.Errorf("RedisClient DecrCtx error: %v", err)
		return
	}

	resp = new(pb.DelCollectionResponse)
	err = nil

	return
}
