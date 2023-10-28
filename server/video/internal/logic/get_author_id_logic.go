package logic

import (
	"context"

	"douniu/server/video/internal/svc"
	"douniu/server/video/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAuthorIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAuthorIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAuthorIdLogic {
	return &GetAuthorIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAuthorIdLogic) GetAuthorId(in *pb.GetAuthorIdReq) (resp *pb.GetAuthorIdResp, err error) {
	// todo: add your logic here and delete this line

	resp = new(pb.GetAuthorIdResp)

	return
}
