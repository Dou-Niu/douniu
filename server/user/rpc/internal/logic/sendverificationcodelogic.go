package logic

import (
	"context"

	"douniu/server/user/rpc/internal/svc"
	"douniu/server/user/rpc/types/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendVerificationCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendVerificationCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendVerificationCodeLogic {
	return &SendVerificationCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 发送验证码
func (l *SendVerificationCodeLogic) SendVerificationCode(in *pb.SendVerificationCodeReq) (*pb.SendVerificationCodeResp, error) {
	// todo: add your logic here and delete this line

	return &pb.SendVerificationCodeResp{}, nil
}
