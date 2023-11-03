package logic

import (
	"context"
	"douniu/server/common/consts"
	"douniu/server/common/errorx"
	"douniu/server/common/utils"
	"douniu/server/user/rpc/types/pb"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"

	"douniu/server/user/api/internal/svc"
	"douniu/server/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChangePasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChangePasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChangePasswordLogic {
	return &ChangePasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChangePasswordLogic) ChangePassword(req *types.ChangePasswordReq) error {
	err := utils.DefaultGetValidParams(l.ctx, req)
	if err != nil {
		return errors.Wrapf(errorx.NewCodeError(1, fmt.Sprintf("validate校验错误: %v", err)), "validate校验错误err :%v", err)
	}
	meId, err := l.ctx.Value(consts.UserId).(json.Number).Int64()

	_, err = l.svcCtx.UserRpc.ChangePassword(l.ctx, &pb.ResetPassword{
		UserId:      meId,
		NewPassword: req.NewPassword,
	})
	if err != nil {
		return errors.Wrapf(err, "req: %+v", req)
	}

	return nil
}
