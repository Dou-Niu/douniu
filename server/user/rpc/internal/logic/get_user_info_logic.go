package logic

import (
	"context"

	"douniu/server/user/rpc/internal/svc"
	"douniu/server/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取用户信息
func (l *GetUserInfoLogic) GetUserInfo(in *pb.UserInfoReq) (resp *pb.UserInfoResp, err error) {
	// todo: add your logic here and delete this line

	resp = new(pb.UserInfoResp)

	return
}
