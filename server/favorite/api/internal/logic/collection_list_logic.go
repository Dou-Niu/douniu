package logic

import (
	"context"
	"douniu/server/common/consts"
	"douniu/server/favorite/rpc/favoriterpc"
	"douniu/server/video/rpc/videorpc"
	"encoding/json"
	"github.com/jinzhu/copier"

	"douniu/server/favorite/api/internal/svc"
	"douniu/server/favorite/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CollectionListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCollectionListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CollectionListLogic {
	return &CollectionListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CollectionListLogic) CollectionList(req *types.CollectioneListRequest) (resp *types.CollectioneListResponse, err error) {
	userId, _ := l.ctx.Value(consts.UserId).(json.Number).Int64()

	videoIdListResp, err := l.svcCtx.FavoriteRpc.GetUserCollectionIdList(l.ctx, &favoriterpc.GetUserCollectionIdListRequest{
		UserId:  req.UserId,
		PageNum: req.PageNum,
	})
	if err != nil {
		l.Errorf("FavoriteRpc GetUserCollectionList error: %v", err)
		return
	}

	listResp, err := l.svcCtx.VideoRpc.GetVideoListInfo(l.ctx, &videorpc.GetVideoListInfoReq{
		MeUserId:    userId,
		VideoIdList: videoIdListResp.VideoIdList,
	})
	if err != nil {
		l.Errorf("VideoRpc GetVideoListInfo error: %v", err)
		return
	}

	resp = new(types.CollectioneListResponse)
	resp.VideoList = make([]*types.Video, 0, len(listResp.VideoList))
	_ = copier.Copy(&resp.VideoList, &listResp.VideoList)

	return
}
