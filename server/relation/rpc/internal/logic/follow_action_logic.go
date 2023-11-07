package logic

import (
	"context"
	consts2 "douniu/common/consts"
	"douniu/server/chat/rpc/chatrpc"
	"douniu/server/relation/model"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"strconv"
	"time"

	"douniu/server/relation/rpc/internal/svc"
	"douniu/server/relation/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type FollowActionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFollowActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowActionLogic {
	return &FollowActionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FollowActionLogic) FollowAction(in *pb.FollowActionRequest) (resp *pb.FollowActionResponse, err error) {
	userIdStr := strconv.FormatInt(in.UserId, 10)
	ToUserIdStr := strconv.FormatInt(in.ToUserId, 10)
	if in.ActionType == consts2.FollowAdd {
		// 判断是否已经关注
		isFollow, err := l.svcCtx.RedisClient.ZscoreCtx(l.ctx, consts2.UserFollowPrefix+userIdStr, ToUserIdStr)
		if err != nil && !errors.Is(err, redis.Nil) {
			l.Errorf("redis ZscoreCtx err: %v", err)
			return nil, err
		}
		if isFollow != 0 {
			return nil, errors.New("您已经关注了该用户")
		}
		// 加入关注列表
		_, err = l.svcCtx.RedisClient.ZaddCtx(l.ctx, consts2.UserFollowPrefix+userIdStr, time.Now().Unix(), ToUserIdStr)
		if err != nil {
			l.Errorf("redis zadd err: %v", err)
			return nil, err
		}
		// 加入粉丝列表
		_, err = l.svcCtx.RedisClient.ZaddCtx(l.ctx, consts2.UserFollowerPrefix+ToUserIdStr, time.Now().Unix(), userIdStr)
		if err != nil {
			l.Errorf("redis zadd err: %v", err)
			return nil, err
		}
		// 判断对方是否是自己的粉丝
		isFollower, err := l.svcCtx.RedisClient.ZscoreCtx(l.ctx, consts2.UserFollowerPrefix+userIdStr, ToUserIdStr)
		if err != nil && !errors.Is(err, redis.Nil) {
			l.Errorf("redis ZscoreCtx err: %v", err)
			return nil, err
		}
		if isFollower != 0 {
			// 发消息
			_, err = l.svcCtx.ChatRpc.MessageAction(l.ctx, &chatrpc.MessageActionRequest{
				FromUserId: in.UserId,
				ToUserId:   in.ToUserId,
				Content:    "我们已经是好友了，快来聊天吧！",
				Action:     consts2.MessageSend,
			})
			if err != nil {
				l.Errorf("chatrpc message action err: %v", err)
				return nil, err
			}
		}

	} else {
		// 判断是否已经关注
		isFollow, err := l.svcCtx.RedisClient.ZscoreCtx(l.ctx, consts2.UserFollowPrefix+userIdStr, ToUserIdStr)
		if err != nil && !errors.Is(err, redis.Nil) {
			l.Errorf("redis ZscoreCtx err: %v", err)
			return nil, err
		}
		if isFollow == 0 {
			return nil, errors.New("您还没有关注该用户")
		}

		// 移除关注列表
		_, err = l.svcCtx.RedisClient.ZremCtx(l.ctx, consts2.UserFollowPrefix+userIdStr, ToUserIdStr)
		if err != nil {
			l.Errorf("redis zrem err: %v", err)
			return nil, err
		}
		// 移除粉丝列表
		_, err = l.svcCtx.RedisClient.ZremCtx(l.ctx, consts2.UserFollowerPrefix+ToUserIdStr, userIdStr)
		if err != nil {
			l.Errorf("redis zrem err: %v", err)
			return nil, err
		}
	}

	// 丢到kafka去落库
	//inStr, err := jsonx.MarshalToString(in)
	//if err != nil {
	//	return nil, err
	//}
	//err = l.svcCtx.KafkaPusher.Push(inStr)
	//if err != nil {
	//	l.Errorf("kafka push err: %v", err)
	//	return nil, err
	//}
	if in.ActionType == consts2.FollowAdd {
		_, err = l.svcCtx.FollowModel.Insert(context.Background(), &model.Follow{
			Id:         l.svcCtx.Snowflake.Generate().Int64(),
			UserId:     in.UserId,
			FollowId:   in.ToUserId,
			UpdateTime: time.Now().Unix(),
		})
		if err != nil {
			logx.Errorf("FollowModel.Insert err: %v", err)
			return
		}
	} else {
		err = l.svcCtx.FollowModel.DeleteByUIdAndFId(context.Background(), in.UserId, in.ToUserId)
		if err != nil {
			logx.Errorf("FollowModel.Delete err: %v", err)
			return
		}
	}

	resp = new(pb.FollowActionResponse)

	return
}
