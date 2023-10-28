package logic

import (
	"context"

	"douniu/server/user/rpc/internal/svc"
	"douniu/server/user/rpc/types/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterOrLoginByPhoneLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterOrLoginByPhoneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterOrLoginByPhoneLogic {
	return &RegisterOrLoginByPhoneLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 使用验证码进行手机号注册或登录
func (l *RegisterOrLoginByPhoneLogic) RegisterOrLoginByPhone(in *pb.RegisterOrLoginByPhoneReq) (*pb.RegisterOrLoginResp, error) {
	// todo: add your logic here and delete this line

	return &pb.RegisterOrLoginResp{}, nil
}
