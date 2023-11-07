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

type RelationFollowerListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRelationFollowerListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RelationFollowerListLogic {
	return &RelationFollowerListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RelationFollowerListLogic) RelationFollowerList(req *types.RelationFollowerListRequest) (resp *types.RelationFollowerListResponse, err error) {
	userId, _ := l.ctx.Value(consts.UserId).(json.Number).Int64()

	resp = new(types.RelationFollowerListResponse)
	followerListResp, err := l.svcCtx.RelationRpc.GetFollowerList(l.ctx, &relationrpc.GetFollowerListRequest{
		UserId:   userId,
		ToUserId: req.UserId,
		PageNum:  req.PageNum,
	})
	if err != nil {
		return
	}
	resp.UserList = make([]*types.User, 0, len(followerListResp.UserList))
	_ = copier.Copy(&resp.UserList, &followerListResp.UserList)

	return
}
