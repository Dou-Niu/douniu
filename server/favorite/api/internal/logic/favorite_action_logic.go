package logic

import (
	"context"

	"douniu/server/favorite/api/internal/svc"
	"douniu/server/favorite/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavoriteActionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFavoriteActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteActionLogic {
	return &FavoriteActionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FavoriteActionLogic) FavoriteAction(req *types.FavoriteLikeRequest) (resp *types.FavoriteLikeResponse, err error) {
	resp = new(types.FavoriteLikeResponse)

	return
}
