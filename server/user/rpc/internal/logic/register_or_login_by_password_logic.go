package logic

import (
	"context"
	"douniu/common/errorx"
	"douniu/common/utils"
	"github.com/pkg/errors"
	"time"

	"douniu/server/user/rpc/internal/svc"
	"douniu/server/user/rpc/types/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterOrLoginByPasswordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterOrLoginByPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterOrLoginByPasswordLogic {
	return &RegisterOrLoginByPasswordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 使用密码进行手机号登录
func (l *RegisterOrLoginByPasswordLogic) RegisterOrLoginByPassword(in *pb.RegisterOrLoginByPasswordReq) (*pb.RegisterOrLoginResp, error) {
	u, err := l.svcCtx.UserModel.FindOneByPhone(l.ctx, in.Phone)
	if err != nil {
		return nil, errors.Wrapf(errorx.NewDefaultError("登录失败，手机号或者密码错误"), "注册失败，userinfo 写入mysql错误 RegisterReq：%v", in)
	}
	if utils.Md5Password(in.Password, l.svcCtx.Config.Salt) != u.Password {
		return nil, errors.Wrapf(errorx.NewDefaultError("登录失败，手机号或者密码错误"), "注册失败，userinfo 写入mysql错误 RegisterReq：%v", in)

	}
	userId := u.Id
	auth := l.svcCtx.Config.JWTAuth
	now := time.Now().Unix()
	accessToken, _ := getJwtToken(auth.AccessSecret, now, auth.AccessTokenExpire, userId)
	// 生成refreshToken
	refreshToken, _ := getJwtToken(auth.AccessSecret, now, auth.RefreshTokenExpire, userId)
	return &pb.RegisterOrLoginResp{
		UserId:       userId,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil

}
