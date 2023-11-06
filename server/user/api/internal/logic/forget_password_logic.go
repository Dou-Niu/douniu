package logic

import (
	"context"
	"douniu/common/errorx"
	"douniu/common/utils"
	"douniu/server/user/rpc/types/pb"
	"fmt"
	"github.com/pkg/errors"

	"douniu/server/user/api/internal/svc"
	"douniu/server/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ForgetPasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewForgetPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ForgetPasswordLogic {
	return &ForgetPasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ForgetPasswordLogic) ForgetPassword(req *types.ForgetPasswordReq) error {
	err := utils.DefaultGetValidParams(l.ctx, req)
	if err != nil {
		return errors.Wrapf(errorx.NewCodeError(1, fmt.Sprintf("validate校验错误: %v", err)), "validate校验错误err :%v", err)
	}
	_, err = l.svcCtx.UserRpc.ForgetPassword(l.ctx, &pb.ForgetPasswordReq{
		NewPassword: req.NewPassword,
		Phone:       req.Phone,
	})
	if err != nil {
		return errors.Wrapf(err, "req: %+v", req)
	}

	return nil

}
