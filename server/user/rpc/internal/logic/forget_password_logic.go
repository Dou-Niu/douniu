package logic

import (
	"context"
	"douniu/common/errorx"
	"douniu/common/utils"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logc"

	"douniu/server/user/rpc/internal/svc"
	"douniu/server/user/rpc/types/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ForgetPasswordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewForgetPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ForgetPasswordLogic {
	return &ForgetPasswordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 忘记密码并重置密码
func (l *ForgetPasswordLogic) ForgetPassword(in *pb.ForgetPasswordReq) (*pb.CommonResp, error) {
	u, err := l.svcCtx.UserModel.FindOneByPhone(l.ctx, in.Phone)
	if err != nil {
		return nil, errors.Wrapf(errorx.NewDefaultError("查询用户信息失败,err:"+err.Error()), "redis查询手机号是否已经注册时候失败 RegisterOrLoginByPhoneReq：%v", in)
	}
	err = l.svcCtx.UserModel.ChangePassword(l.ctx, u.Id, utils.Md5Password(in.NewPassword, l.svcCtx.Config.Salt))
	if err != nil {
		logc.Error(l.ctx, err)
		return nil, errors.Wrapf(errorx.NewDefaultError("更改用户密码失败"+err.Error()), "更改用户密码失败ResetPassword：%v", in)

	}

	return &pb.CommonResp{}, nil
}
