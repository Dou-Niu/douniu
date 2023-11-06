package logic

import (
	"context"
	"douniu/common/consts"
	"douniu/server/favorite/api/internal/svc"
	"douniu/server/favorite/api/internal/types"
	"douniu/server/favorite/rpc/favoriterpc"
	"encoding/json"

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
	userId, err := l.ctx.Value(consts.UserId).(json.Number).Int64()

	if req.ActionType == consts.FavoriteAdd {
		_, err = l.svcCtx.FavoriteRpc.AddFavorite(l.ctx, &favoriterpc.AddFavoriteRequest{
			UserId:  userId,
			VideoId: req.VideoId,
		})
		if err != nil {
			l.Errorf("AddFavorite error: %v", err)
			return
		}

	} else {
		_, err = l.svcCtx.FavoriteRpc.DelFavorite(l.ctx, &favoriterpc.DelFavoriteRequest{
			UserId:  userId,
			VideoId: req.VideoId,
		})
		if err != nil {
			l.Errorf("DelFavorite error: %v", err)
			return
		}
	}

	return
}
