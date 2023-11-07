package logic

import (
	"context"
	consts2 "douniu/common/consts"
	"douniu/common/errorx"
	"fmt"
	"github.com/pkg/errors"

	"douniu/server/video/rpc/internal/svc"
	"douniu/server/video/rpc/types/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type FeedUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	GetVideoListInfo *GetVideoListInfoLogic
}

func NewFeedUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FeedUserLogic {
	return &FeedUserLogic{
		ctx:              ctx,
		svcCtx:           svcCtx,
		Logger:           logx.WithContext(ctx),
		GetVideoListInfo: NewGetVideoListInfoLogic(ctx, svcCtx),
	}
}

func (l *FeedUserLogic) FeedUser(in *pb.FeedUserReq) (*pb.FeedResp, error) {
	var key string
	// 热度和最新排序
	switch int(in.Sort) {
	case consts2.SortByHot:
		key = consts2.VideoEveryUserHotScore
	case consts2.SortByTime:
		key = consts2.VideoEveryUserTimeScore
	default:
		return nil, errors.Wrapf(errorx.NewDefaultError("视频排序规则sort非法输入"), "视频排序规则sort非法输入 Req：%v", in)

	}
	key += fmt.Sprint(in.UserId)
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
