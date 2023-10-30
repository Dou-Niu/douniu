package logic

import (
	"context"
	"douniu/server/common/consts"
	"douniu/server/video/rpc/internal/svc"
	"douniu/server/video/rpc/types/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type FeedHomeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	GetVideoListInfo *GetVideoListInfoLogic
}

func NewFeedHomeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FeedHomeLogic {
	return &FeedHomeLogic{
		ctx:              ctx,
		svcCtx:           svcCtx,
		Logger:           logx.WithContext(ctx),
		GetVideoListInfo: NewGetVideoListInfoLogic(ctx, svcCtx),
	}
}

func (l *FeedHomeLogic) FeedHome(in *pb.FeedHomeReq) (*pb.FeedHomeResp, error) {
	videoIds, err := l.svcCtx.VideoModel.FindByTimeOrHot(l.ctx, l.svcCtx.RedisClient, consts.VideoTimeScore, in.LatestTime)
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
	return &pb.FeedHomeResp{
		VideoList: list,
		NextTime:  list[len(list)-1].CreateTime,
	}, nil
}
