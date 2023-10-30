package logic

import (
	"context"
	"douniu/server/common/consts"
	"douniu/server/common/errorx"
	"fmt"
	"github.com/pkg/errors"

	"douniu/server/video/rpc/internal/svc"
	"douniu/server/video/rpc/types/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type FeedPartitionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	GetVideoListInfo *GetVideoListInfoLogic
}

func NewFeedPartitionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FeedPartitionLogic {
	return &FeedPartitionLogic{
		ctx:              ctx,
		svcCtx:           svcCtx,
		Logger:           logx.WithContext(ctx),
		GetVideoListInfo: NewGetVideoListInfoLogic(ctx, svcCtx),
	}
}

func (l *FeedPartitionLogic) FeedPartition(in *pb.FeedPartitionReq) (*pb.FeedResp, error) {
	var key string
	// 热度和最新排序
	switch int(in.Sort) {
	case consts.SortByHot:
		key = consts.VideoPartitionHotScore
	case consts.SortByTime:
		key = consts.VideoPartitionTimeScore
	default:
		return nil, errors.Wrapf(errorx.NewDefaultError("视频排序规则sort非法输入"), "视频排序规则sort非法输入 Req：%v", in)

	}
	key += fmt.Sprint(in.Partition)
	videoIds, err := l.svcCtx.VideoModel.FindByTimeOrHot(l.ctx, l.svcCtx.RedisClient, key, in.MaxValue)
	if err != nil {
		return nil, err
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
