package logic

import (
	"context"
	consts2 "douniu/common/consts"
	"douniu/server/video/rpc/videorpc"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/mr"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"strconv"
	"time"

	"douniu/server/favorite/rpc/internal/svc"
	"douniu/server/favorite/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelFavoriteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelFavoriteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelFavoriteLogic {
	return &DelFavoriteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelFavoriteLogic) DelFavorite(in *pb.DelFavoriteRequest) (resp *pb.DelFavoriteResponse, err error) {
	userIdStr := strconv.Itoa(int(in.UserId))
	videoIdStr := strconv.Itoa(int(in.VideoId))
	partitionIdStr := strconv.Itoa(int(in.Partition))
	// 判断是否已经点赞
	isFavorited, err := l.svcCtx.RedisClient.ZscoreCtx(l.ctx, consts2.UserCollectPrefix+userIdStr, videoIdStr)
	if err != nil && !errors.Is(err, redis.Nil) {
		l.Errorf("RedisClient ZscoreCtx error: %v", err)
		return
	}
	if isFavorited == 0 {
		return nil, errors.New("你还没有点赞过")
	}

	authorIdResp, err := l.svcCtx.VideoRpc.GetAuthorId(l.ctx, &videorpc.GetAuthorIdReq{VideoId: in.VideoId})
	if err != nil {
		l.Errorf("VideoRpc GetAuthorId error: %v", err)
		return
	}
	authorId := authorIdResp.UserId

	err = mr.Finish(func() error {
		// 视频移除到用户的点赞列表
		_, err = l.svcCtx.RedisClient.ZremCtx(l.ctx, consts2.UserFavoriteIdPrefix+userIdStr, videoIdStr)
		if err != nil {
			l.Errorf("RedisClient ZremCtx error: %v", err)
			return err
		}
		return nil
	}, func() error {
		// 用户移除到视频的点赞列表
		_, err = l.svcCtx.RedisClient.ZremCtx(l.ctx, consts2.VideoFavoritedIdPrefix+videoIdStr, userIdStr)
		if err != nil {
			l.Errorf("RedisClient ZremCtx error: %v", err)
			return err
		}
		return nil
	}, func() error {
		// 作者的获赞数-1
		_, err = l.svcCtx.RedisClient.DecrCtx(l.ctx, consts2.UserFavoritedCountPrefix+strconv.Itoa(int(authorId)))
		if err != nil {
			l.Errorf("RedisClient IncrCtx error: %v", err)
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

	resp = new(pb.DelFavoriteResponse)
	err = nil

	// 查询点赞记录
	favorite, err := l.svcCtx.FavoriteModel.FindByUserIdVideoId(l.ctx, in.UserId, in.VideoId)
	if err != nil {
		l.Errorf("FavoriteModel FindByUserIdVideoId error: %v", err)
		return
	}

	// 更新点赞记录
	if favorite != nil {
		favorite.Status = 2
		favorite.UpdateTime = time.Now().Unix()
		err := l.svcCtx.FavoriteModel.Update(l.ctx, favorite)
		if err != nil {
			l.Errorf("FavoriteModel Update error: %v", err)
		}

	}

	return
}
