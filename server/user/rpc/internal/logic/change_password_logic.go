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
	err := l.svcCtx.UserModel.ChangePassword(l.ctx, in.UserId, utils.Md5Password(in.NewPassword, l.svcCtx.Config.Salt))
	if err != nil {
		logc.Error(l.ctx, err)
		return nil, errors.Wrapf(errorx.NewDefaultError("更改用户密码失败"+err.Error()), "更改用户密码失败ResetPassword：%v", in)

	}
	return &pb.CommonResp{}, nil
}
