package logic

import (
	"context"
	"douniu/server/common/consts"

	"douniu/server/video/rpc/internal/svc"
	"douniu/server/video/rpc/types/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type FeedHotLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	GetVideoListInfo *GetVideoListInfoLogic
}

func NewFeedHotLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FeedHotLogic {
	return &FeedHotLogic{
		ctx:              ctx,
		svcCtx:           svcCtx,
		Logger:           logx.WithContext(ctx),
		GetVideoListInfo: NewGetVideoListInfoLogic(ctx, svcCtx),
	}
}

func (l *FeedHotLogic) FeedHot(in *pb.FeedHotReq) (*pb.FeedHotResp, error) {
	videoIds, err := l.svcCtx.VideoModel.FindByTimeOrHot(l.ctx, l.svcCtx.RedisClient, consts.VideoHotScore, in.MaxHot)
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

	return &pb.FeedHotResp{
		NextMaxHost: list[len(list)-1].CreateTime,
		VideoList:   list,
	}, nil
}
