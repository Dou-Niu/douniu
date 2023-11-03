package logic

import (
	"context"
	"douniu/server/common/errorx"
	"fmt"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logc"
	"time"

	"douniu/server/video/rpc/internal/svc"
	"douniu/server/video/rpc/types/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type FeedFollowLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	GetVideoListInfo *GetVideoListInfoLogic
}

func NewFeedFollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FeedFollowLogic {
	return &FeedFollowLogic{
		ctx:              ctx,
		svcCtx:           svcCtx,
		Logger:           logx.WithContext(ctx),
		GetVideoListInfo: NewGetVideoListInfoLogic(ctx, svcCtx),
	}
}

func (l *FeedFollowLogic) FeedFollow(in *pb.FeedFollowReq) (*pb.FeedResp, error) {
	// TODO 调用获取用户关注列表id的rpc
	followIds := []int64{1, 2}
	videoIds, err := l.svcCtx.VideoModel.FindFollowFeed(l.ctx, followIds, time.Unix(in.LatestTime, 0))
	if err != nil {
		logc.Error(l.ctx, fmt.Sprintf("根据follow_id查询video_id错误,err:%v", err))
		return nil, errors.Wrapf(errorx.NewDefaultError(err.Error()), "根据follow_id查询video_id错误,err")
	}
	videoList, err := l.GetVideoListInfo.GetVideoListInfo(&pb.GetVideoListInfoReq{
		MeUserId:    in.MeUserId,
		VideoIdList: videoIds,
	})
	if err != nil {
		return nil, err
	}
	list := videoList.VideoList
	return &pb.FeedResp{
		VideoList:    list,
		NextMaxValue: list[len(list)-1].CreateTime,
	}, nil
}
