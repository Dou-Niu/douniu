package logic

import (
	"context"

	"douniu/server/user/rpc/internal/svc"
	"douniu/server/user/rpc/types/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChangePasswordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewChangePasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChangePasswordLogic {
	return &ChangePasswordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 修改密码
func (l *ChangePasswordLogic) ChangePassword(in *pb.ResetPassword) (*pb.CommonResp, error) {
	// todo: add your logic here and delete this line

	return &pb.CommonResp{}, nil
}
