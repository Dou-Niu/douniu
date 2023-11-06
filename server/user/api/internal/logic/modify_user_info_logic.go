package logic

import (
	"context"
	"douniu/common/consts"
	"douniu/common/errorx"
	"douniu/common/utils"
	"douniu/server/user/api/internal/svc"
	"douniu/server/user/api/internal/types"
	"douniu/server/user/rpc/types/pb"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type ModifyUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewModifyUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ModifyUserInfoLogic {
	return &ModifyUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ModifyUserInfoLogic) ModifyUserInfo(req *types.ModifyUserInfoReq) error {
	err := utils.DefaultGetValidParams(l.ctx, req)
	if err != nil {
		return errors.Wrapf(errorx.NewCodeError(1, fmt.Sprintf("validate校验错误: %v", err)), "validate校验错误err :%v", err)
	}

	meId, err := l.ctx.Value(consts.UserId).(json.Number).Int64()
	_, err = l.svcCtx.UserRpc.ModifyUserInfo(l.ctx, &pb.ModifyUserInfoReq{
		MeUserId: meId,
		Types:    req.Types,
		Value:    req.Value,
	})
	if err != nil {
		return errors.Wrapf(err, "req: %+v", req)
	}
	return nil
}
