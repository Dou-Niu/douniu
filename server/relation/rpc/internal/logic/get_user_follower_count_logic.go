package logic

import (
	"context"

	"douniu/server/relation/rpc/internal/svc"
	"douniu/server/relation/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserFollowerCountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserFollowerCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserFollowerCountLogic {
	return &GetUserFollowerCountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserFollowerCountLogic) GetUserFollowerCount(in *pb.GetUserFollowerCountRequest) (resp *pb.GetUserFollowerCountResponse, err error) {
	// todo: add your logic here and delete this line

	resp = new(pb.GetUserFollowerCountResponse)

	return
}
