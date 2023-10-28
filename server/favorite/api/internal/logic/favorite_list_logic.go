package logic

import (
	"context"
	"douniu/server/common/consts"
	"douniu/server/favorite/rpc/favoriterpc"
	"douniu/server/video/videorpc"
	"encoding/json"
	"github.com/jinzhu/copier"

	"douniu/server/favorite/api/internal/svc"
	"douniu/server/favorite/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavoriteListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFavoriteListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteListLogic {
	return &FavoriteListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FavoriteListLogic) FavoriteList(req *types.FavoriteListRequest) (resp *types.FavoriteListResponse, err error) {
	userId, _ := l.ctx.Value(consts.UserId).(json.Number).Int64()

	videoIdListResp, err := l.svcCtx.FavoriteRpc.GetFavoriteVideoIdList(l.ctx, &favoriterpc.GetFavoriteVideoIdListRequest{
		UserId: req.UserId,
	})
	if err != nil {
		l.Errorf("FavoriteRpc GetFavoriteVideoIdList error: %v", err)
		return
	}
	listResp, err := l.svcCtx.VideoRpc.GetVideoListInfo(l.ctx, &videorpc.GetVideoListInfoReq{
		UserId:      userId,
		VideoIdList: videoIdListResp.VideoIdList,
	})
	if err != nil {
		l.Errorf("VideoRpc GetVideoListInfo error: %v", err)
		return
	}

	resp = new(types.FavoriteListResponse)
	resp.VideoList = make([]*types.Video, 0, len(listResp.VideoList))
	_ = copier.Copy(&resp.VideoList, &listResp.VideoList)

	return
}
