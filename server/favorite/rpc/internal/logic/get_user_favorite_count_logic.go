package logic

import (
	"context"

	"douniu/server/favorite/rpc/internal/svc"
	"douniu/server/favorite/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserFavoriteCountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserFavoriteCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserFavoriteCountLogic {
	return &GetUserFavoriteCountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserFavoriteCountLogic) GetUserFavoriteCount(in *pb.GetUserFavoriteCountRequest) (resp *pb.GetUserFavoriteCountResponse, err error) {
	// todo: add your logic here and delete this line

	resp = new(pb.GetUserFavoriteCountResponse)

	return
}
