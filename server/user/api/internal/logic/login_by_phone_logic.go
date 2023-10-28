package logic

import (
	"context"
	"douniu/server/common/errorx"
	"douniu/server/common/utils"
	"douniu/server/user/rpc/types/pb"
	"fmt"
	"github.com/pkg/errors"

	"douniu/server/user/api/internal/svc"
	"douniu/server/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginByPhoneLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginByPhoneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginByPhoneLogic {
	return &LoginByPhoneLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginByPhoneLogic) LoginByPhone(req *types.RegisterOrLoginByPhoneReq) (resp *types.RegisterOrLoginResp, err error) {
	err = utils.DefaultGetValidParams(l.ctx, req)
	if err != nil {
		return nil, errors.Wrapf(errorx.NewCodeError(1, fmt.Sprintf("validate校验错误: %v", err)), "validate校验错误err :%v", err)
	}
	res, err := l.svcCtx.UserRpc.RegisterOrLoginByPhone(l.ctx, &pb.RegisterOrLoginByPhoneReq{
		Phone:            req.Phone,
		VerificationCode: req.VerificationCode,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}
	return &types.RegisterOrLoginResp{
		UserId:       res.UserId,
		AccessToken:  res.AccessToken,
		RefreshToken: res.RefreshToken,
	}, nil

}
