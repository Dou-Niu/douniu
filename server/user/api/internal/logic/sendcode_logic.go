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

type SendcodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendcodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendcodeLogic {
	return &SendcodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendcodeLogic) Sendcode(req *types.SendVerificationCodeReq) (resp *types.SendVerificationCodeResp, err error) {
	err = utils.DefaultGetValidParams(l.ctx, req)
	if err != nil {
		return nil, errors.Wrapf(errorx.NewCodeError(1, fmt.Sprintf("validate校验错误: %v", err)), "validate校验错误err :%v", err)
	}
	res, err := l.svcCtx.UserRpc.SendVerificationCode(l.ctx, &pb.SendVerificationCodeReq{Phone: req.Phone})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}
	return &types.SendVerificationCodeResp{
		VerificationCode: res.VerificationCode,
	}, nil

}
