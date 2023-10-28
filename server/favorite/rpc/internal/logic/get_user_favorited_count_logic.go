package logic

import (
	"context"

	"douniu/server/favorite/rpc/internal/svc"
	"douniu/server/favorite/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserFavoritedCountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserFavoritedCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserFavoritedCountLogic {
	return &GetUserFavoritedCountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserFavoritedCountLogic) GetUserFavoritedCount(in *pb.GetUserFavoritedCountRequest) (resp *pb.GetUserFavoritedCountResponse, err error) {
	// todo: add your logic here and delete this line

	resp = new(pb.GetUserFavoritedCountResponse)

	return
}
