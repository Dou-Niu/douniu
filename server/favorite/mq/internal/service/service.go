package service

import (
	"context"
	"douniu/server/common/consts"
	"douniu/server/favorite/model"
	"douniu/server/favorite/mq/internal/config"
	"douniu/server/favorite/rpc/pb"
	"douniu/server/video/rpc/videorpc"
	"github.com/bwmarrin/snowflake"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"time"
)

type Service struct {
	Config        config.Config
	FavoriteModel model.FavoriteModel
	RedisClient   *redis.Redis
	VideoRpc      videorpc.VideoRpc
	Snowflake     *snowflake.Node
}

func NewService(c config.Config) *Service {
	snowflakeNode, _ := snowflake.NewNode(consts.FavoriteMachineId)
	return &Service{
		Config:        c,
		FavoriteModel: model.NewFavoriteModel(sqlx.NewMysql(c.MySQLConf.DataSource)),
		RedisClient:   redis.MustNewRedis(c.RedisConf),
		Snowflake:     snowflakeNode,
		VideoRpc:      videorpc.NewVideoRpc(zrpc.MustNewClient(c.VideoRpcConf)),
	}
}

func (s *Service) Consume(_ string, value string) error {
	logx.Info("成功消费消息", value)
	var in pb.AddFavoriteRequest
	err := jsonx.UnmarshalFromString(value, &in)
	if err != nil {
		logx.Errorf("UnmarshalFromString，err：%v", err)
		return err
	}

	authorIdResp, err := s.VideoRpc.GetAuthorId(context.Background(), &videorpc.GetAuthorIdReq{VideoId: in.VideoId})
	if err != nil {
		logx.Errorf("VideoRpc GetAuthorId error: %v", err)
		return err
	}
	authorId := authorIdResp.UserId

	// 查询点赞记录
	favorite, err := s.FavoriteModel.FindByUserIdVideoId(context.Background(), in.UserId, in.VideoId)
	if err != nil {
		logx.Errorf("FavoriteModel FindByUserIdVideoId error: %v", err)
		return err
	}

	// 创建或更新点赞记录
	if favorite != nil {
		favorite.Status = 1
		favorite.UpdateTime = time.Now().Unix()
		err := s.FavoriteModel.Update(context.Background(), favorite)
		if err != nil {
			logx.Errorf("FavoriteModel Update error: %v", err)
		}
		return err

	} else {
		_, err := s.FavoriteModel.Insert(context.Background(), &model.Favorite{
			Id:         s.Snowflake.Generate().Int64(),
			UserId:     in.UserId,
			VideoId:    in.VideoId,
			AuthorId:   authorId,
			Status:     1,
			UpdateTime: time.Now().Unix(),
		})
		if err != nil {
			logx.Errorf("FavoriteModel Insert error: %v", err)
		}
		return err
	}

}
