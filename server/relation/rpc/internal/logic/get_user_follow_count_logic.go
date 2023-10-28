package logic

import (
	"context"

	"douniu/server/relation/rpc/internal/svc"
	"douniu/server/relation/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserFollowCountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserFollowCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserFollowCountLogic {
	return &GetUserFollowCountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserFollowCountLogic) GetUserFollowCount(in *pb.GetUserFollowCountRequest) (resp *pb.GetUserFollowCountResponse, err error) {
	// todo: add your logic here and delete this line

	resp = new(pb.GetUserFollowCountResponse)

	return
}
