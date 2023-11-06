package logic

import (
	"context"
	consts2 "douniu/common/consts"
	"errors"
	"github.com/zeromicro/go-zero/core/mr"
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
	partitionIdStr := strconv.Itoa(int(in.Partition))
	// 判断是否已经收藏
	isFavorited, err := l.svcCtx.RedisClient.ZscoreCtx(l.ctx, consts2.UserCollectPrefix+userIdStr, videoIdStr)
	if err != nil && !errors.Is(err, redis.Nil) {
		l.Errorf("RedisClient ZscoreCtx error: %v", err)
		return
	}
	if isFavorited != 0 {
		return nil, errors.New("你已经收藏过了")
	}

	err = mr.Finish(func() error {
		// 视频添加到用户的点收藏列表
		_, err = l.svcCtx.RedisClient.ZaddCtx(l.ctx, consts2.UserCollectPrefix+userIdStr, time.Now().Unix(), videoIdStr)
		if err != nil {
			l.Errorf("RedisClient ZaddCtx error: %v", err)
			return err
		}
		return nil
	}, func() error {
		// 视频收藏数量+1
		_, err = l.svcCtx.RedisClient.IncrbyCtx(l.ctx, consts2.VideoCollectCountPrefix+videoIdStr, 1)
		if err != nil {
			l.Errorf("RedisClient IncrbyCtx error: %v", err)
			return err
		}
		return nil
	}, func() error {
		// 视频热度增加
		_, err = l.svcCtx.RedisClient.ZincrbyCtx(l.ctx, consts2.VideoHotScore, int64(consts2.SingleHotScore), videoIdStr)
		if err != nil {
			l.Errorf("RedisClient ZincrbyCtx error: %v", err)
			return err
		}
		return nil
	}, func() error {
		// 视频分区热度增加
		_, err = l.svcCtx.RedisClient.ZincrbyCtx(l.ctx, consts2.VideoPartitionHotScore+partitionIdStr, int64(consts2.SingleHotScore), videoIdStr)
		if err != nil {
			l.Errorf("RedisClient ZincrbyCtx error: %v", err)
			return err
		}
		return nil
	}, func() error {
		// 用户视频热度增加
		_, err = l.svcCtx.RedisClient.ZincrbyCtx(l.ctx, consts2.VideoEveryUserHotScore+userIdStr, int64(consts2.SingleHotScore), videoIdStr)
		if err != nil {
			l.Errorf("RedisClient ZincrbyCtx error: %v", err)
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	resp = new(pb.AddCollectionResponse)
	err = nil

	return
}
