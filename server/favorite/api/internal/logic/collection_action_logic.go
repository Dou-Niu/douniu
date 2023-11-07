package logic

import (
	"context"
	"douniu/common/consts"
	"douniu/server/favorite/rpc/favoriterpc"
	"encoding/json"

	"douniu/server/favorite/api/internal/svc"
	"douniu/server/favorite/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CollectionActionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCollectionActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CollectionActionLogic {
	return &CollectionActionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CollectionActionLogic) CollectionAction(req *types.CollectionActionRequest) (resp *types.CollectionActionResponse, err error) {
	userId, err := l.ctx.Value(consts.UserId).(json.Number).Int64()

	if req.ActionType == consts.CollectionAdd {
		_, err = l.svcCtx.FavoriteRpc.AddCollection(l.ctx, &favoriterpc.AddCollectionRequest{
			UserId:    userId,
			VideoId:   req.VideoId,
			Partition: req.Partition,
		})
		if err != nil {
			l.Errorf("AddFavorite error: %v", err)
			return
		}

	} else {
		_, err = l.svcCtx.FavoriteRpc.DelCollection(l.ctx, &favoriterpc.DelCollectionRequest{
			UserId:    userId,
			VideoId:   req.VideoId,
			Partition: req.Partition,
		})
		if err != nil {
			l.Errorf("DelFavorite error: %v", err)
			return
		}
	}

	return
}
