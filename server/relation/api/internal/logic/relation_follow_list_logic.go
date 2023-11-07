package logic

import (
	"context"
	"douniu/common/consts"
	"douniu/server/relation/rpc/relationrpc"
	"encoding/json"
	"github.com/jinzhu/copier"

	"douniu/server/relation/api/internal/svc"
	"douniu/server/relation/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RelationFollowListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRelationFollowListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RelationFollowListLogic {
	return &RelationFollowListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RelationFollowListLogic) RelationFollowList(req *types.RelationFollowListRequest) (resp *types.RelationFollowListResponse, err error) {
	userId, _ := l.ctx.Value(consts.UserId).(json.Number).Int64()

	resp = new(types.RelationFollowListResponse)
	followListResp, err := l.svcCtx.RelationRpc.GetFollowList(l.ctx, &relationrpc.GetFollowListRequest{
		UserId:   userId,
		ToUserId: req.UserId,
		PageNum:  req.PageNum,
	})
	if err != nil {
		return
	}

	resp.UserList = make([]*types.User, 0, len(followListResp.UserList))
	_ = copier.Copy(&resp.UserList, &followListResp.UserList)
	return

}
