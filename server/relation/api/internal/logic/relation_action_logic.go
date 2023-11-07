package logic

import (
	"context"
	"douniu/common/consts"
	"douniu/server/relation/rpc/relationrpc"
	"encoding/json"

	"douniu/server/relation/api/internal/svc"
	"douniu/server/relation/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RelationActionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRelationActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RelationActionLogic {
	return &RelationActionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RelationActionLogic) RelationAction(req *types.RelationActionRequest) (resp *types.RelationActionResponse, err error) {
	userId, _ := l.ctx.Value(consts.UserId).(json.Number).Int64()

	resp = new(types.RelationActionResponse)
	_, err = l.svcCtx.RelationRpc.FollowAction(l.ctx, &relationrpc.FollowActionRequest{
		UserId:     userId,
		ToUserId:   req.ToUserId,
		ActionType: req.ActionType,
	})

	return
}
