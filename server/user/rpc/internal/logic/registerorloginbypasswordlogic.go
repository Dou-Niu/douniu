package logic

import (
	"context"

	"douniu/server/user/rpc/internal/svc"
	"douniu/server/user/rpc/types/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterOrLoginByPasswordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterOrLoginByPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterOrLoginByPasswordLogic {
	return &RegisterOrLoginByPasswordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 使用密码进行手机号注册或登录
func (l *RegisterOrLoginByPasswordLogic) RegisterOrLoginByPassword(in *pb.RegisterOrLoginByPasswordReq) (*pb.RegisterOrLoginResp, error) {

	return &pb.RegisterOrLoginResp{}, nil
}
