package logic

import (
	"context"

	"douniu/server/favorite/rpc/internal/svc"
	"douniu/server/favorite/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFavoriteVideoIdListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFavoriteVideoIdListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFavoriteVideoIdListLogic {
	return &GetFavoriteVideoIdListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFavoriteVideoIdListLogic) GetFavoriteVideoIdList(in *pb.GetFavoriteVideoIdListRequest) (resp *pb.GetFavoriteVideoListIdResponse, err error) {
	// todo: add your logic here and delete this line

	resp = new(pb.GetFavoriteVideoListIdResponse)

	return
}
