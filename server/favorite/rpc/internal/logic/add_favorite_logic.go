package logic

import (
	"context"
	"douniu/server/common/consts"
	"douniu/server/favorite/model"
	"douniu/server/video/rpc/videorpc"
	"github.com/zeromicro/go-zero/core/mr"

	"errors"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"strconv"
	"time"

	"douniu/server/favorite/rpc/internal/svc"
	"douniu/server/favorite/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddFavoriteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddFavoriteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddFavoriteLogic {
	return &AddFavoriteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddFavoriteLogic) AddFavorite(in *pb.AddFavoriteRequest) (resp *pb.AddFavoriteResponse, err error) {
	userIdStr := strconv.Itoa(int(in.UserId))
	videoIdStr := strconv.Itoa(int(in.VideoId))
	partitionIdStr := strconv.Itoa(int(in.Partition))
	// 判断是否已经点赞
	isFavorited, err := l.svcCtx.RedisClient.ZscoreCtx(l.ctx, consts.UserFavoriteIdPrefix+userIdStr, videoIdStr)
	if err != nil && !errors.Is(err, redis.Nil) {
		l.Errorf("RedisClient ZscoreCtx error: %v", err)
		return
	}
	if isFavorited != 0 {
		return nil, errors.New("你已经点赞过了")
	}

	authorIdResp, err := l.svcCtx.VideoRpc.GetAuthorId(l.ctx, &videorpc.GetAuthorIdReq{VideoId: in.VideoId})
	if err != nil {
		l.Errorf("VideoRpc GetAuthorId error: %v", err)
		return
	}
	authorId := authorIdResp.UserId

	err = mr.Finish(func() error {
		// 视频添加到用户的点赞列表
		_, err = l.svcCtx.RedisClient.ZaddCtx(l.ctx, consts.UserFavoriteIdPrefix+userIdStr, time.Now().Unix(), videoIdStr)
		if err != nil {
			l.Errorf("RedisClient ZaddCtx error: %v", err)
			return err
		}
		return nil
	}, func() error {
		// 用户添加到视频的点赞列表
		_, err = l.svcCtx.RedisClient.ZaddCtx(l.ctx, consts.VideoFavoritedIdPrefix+videoIdStr, time.Now().Unix(), userIdStr)
		if err != nil {
			l.Errorf("RedisClient ZaddCtx error: %v", err)
			return err
		}
		return nil
	}, func() error {
		// 作者的获赞数+1
		_, err = l.svcCtx.RedisClient.IncrCtx(l.ctx, consts.UserFavoritedCountPrefix+strconv.Itoa(int(authorId)))
		if err != nil {
			l.Errorf("RedisClient IncrCtx error: %v", err)
			return err
		}
		return nil
	}, func() error {
		// 视频热度增加
		_, err = l.svcCtx.RedisClient.ZincrbyCtx(l.ctx, consts.VideoHotScore, int64(consts.SingleHotScore), videoIdStr)
		if err != nil {
			l.Errorf("RedisClient ZincrbyCtx error: %v", err)
			return err
		}
		return nil
	}, func() error {
		// 视频分区热度增加
		_, err = l.svcCtx.RedisClient.ZincrbyCtx(l.ctx, consts.VideoPartitionHotScore+partitionIdStr, int64(consts.SingleHotScore), videoIdStr)
		if err != nil {
			l.Errorf("RedisClient ZincrbyCtx error: %v", err)
			return err
		}
		return nil
	}, func() error {
		// 用户视频热度增加
		_, err = l.svcCtx.RedisClient.ZincrbyCtx(l.ctx, consts.VideoEveryUserHotScore+userIdStr, int64(consts.SingleHotScore), videoIdStr)
		if err != nil {
			l.Errorf("RedisClient ZincrbyCtx error: %v", err)
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	resp = new(pb.AddFavoriteResponse)
	err = nil

	// 查询点赞记录
	favorite, err := l.svcCtx.FavoriteModel.FindByUserIdVideoId(l.ctx, in.UserId, in.VideoId)
	if err != nil {
		l.Errorf("FavoriteModel FindByUserIdVideoId error: %v", err)
		return
	}

	// 创建或更新点赞记录
	if favorite != nil {
		favorite.Status = 1
		favorite.UpdateTime = time.Now().Unix()
		err := l.svcCtx.FavoriteModel.Update(l.ctx, favorite)
		if err != nil {
			l.Errorf("FavoriteModel Update error: %v", err)
		}

	} else {
		_, err := l.svcCtx.FavoriteModel.Insert(l.ctx, &model.Favorite{
			Id:         l.svcCtx.Snowflake.Generate().Int64(),
			UserId:     in.UserId,
			VideoId:    in.VideoId,
			AuthorId:   authorId,
			Status:     1,
			UpdateTime: time.Now().Unix(),
		})
		if err != nil {
			l.Errorf("FavoriteModel Insert error: %v", err)
		}
	}

	return

}
