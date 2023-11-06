package logic

import (
	"context"
	consts2 "douniu/common/consts"
	"errors"
	"github.com/zeromicro/go-zero/core/mr"
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
	partitionIdStr := strconv.Itoa(int(in.Partition))
	// 判断是否已经收藏
	isFavorited, err := l.svcCtx.RedisClient.ZscoreCtx(l.ctx, consts2.UserCollectPrefix+userIdStr, videoIdStr)
	if err != nil && !errors.Is(err, redis.Nil) {
		l.Errorf("RedisClient ZscoreCtx error: %v", err)
		return
	}
	if isFavorited == 0 {
		return nil, errors.New("你还没有收藏过")
	}

	err = mr.Finish(func() error {
		// 删除收藏
		if _, err = l.svcCtx.RedisClient.ZremCtx(l.ctx, consts2.UserCollectPrefix+userIdStr, videoIdStr); err != nil {
			l.Errorf("RedisClient ZremCtx error: %v", err)
			return err
		}
		return nil
	}, func() error {
		// 视频收藏数-1
		if _, err = l.svcCtx.RedisClient.DecrCtx(l.ctx, consts2.VideoCollectCountPrefix+videoIdStr); err != nil {
			l.Errorf("RedisClient DecrCtx error: %v", err)
			return err
		}
		return nil
	}, func() error {
		// 视频热度减少
		_, err = l.svcCtx.RedisClient.ZincrbyCtx(l.ctx, consts2.VideoHotScore, -int64(consts2.SingleHotScore), videoIdStr)
		if err != nil {
			l.Errorf("RedisClient ZincrbyCtx error: %v", err)
			return err
		}
		return nil
	}, func() error {
		// 视频分区热度减少
		_, err = l.svcCtx.RedisClient.ZincrbyCtx(l.ctx, consts2.VideoPartitionHotScore+partitionIdStr, -int64(consts2.SingleHotScore), videoIdStr)
		if err != nil {
			l.Errorf("RedisClient ZincrbyCtx error: %v", err)
			return err
		}
		return nil
	}, func() error {
		// 用户视频热度减少
		_, err = l.svcCtx.RedisClient.ZincrbyCtx(l.ctx, consts2.VideoEveryUserHotScore+userIdStr, -int64(consts2.SingleHotScore), videoIdStr)
		if err != nil {
			l.Errorf("RedisClient ZincrbyCtx error: %v", err)
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	resp = new(pb.DelCollectionResponse)

	return
}
