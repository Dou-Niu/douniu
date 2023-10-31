package logic

import (
	"context"
	"douniu/server/common/errorx"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logc"

	"douniu/server/user/rpc/internal/svc"
	"douniu/server/user/rpc/types/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ModifyUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewModifyUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ModifyUserInfoLogic {
	return &ModifyUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 修改用户信息
func (l *ModifyUserInfoLogic) ModifyUserInfo(in *pb.ModifyUserInfoReq) (*pb.CommonResp, error) {
	err := l.svcCtx.UserModel.ModifyUserInfo(l.ctx, in.MeUserId, in.Types, in.Value)
	if err != nil {
		logc.Error(l.ctx, "修改用户信息失败,err:"+err.Error())
		return nil, errors.Wrapf(errorx.NewDefaultError("修改用户信息失败,err:"+err.Error()), "修改用户信息时候失败 RegisterOrLoginByPhoneReq：%v", in)

	}
	return &pb.CommonResp{}, nil
}
