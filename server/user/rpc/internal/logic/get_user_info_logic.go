package logic

import (
	"context"
	"douniu/server/common/errorx"
	"github.com/pkg/errors"

	"douniu/server/user/rpc/internal/svc"
	"douniu/server/user/rpc/types/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取用户信息
func (l *GetUserInfoLogic) GetUserInfo(in *pb.UserInfoReq) (*pb.UserInfoResp, error) {

	u, err := l.svcCtx.UserModel.FindOne(l.ctx, in.UserId)
	if err != nil {

		return nil, errors.Wrapf(errorx.NewDefaultError("查询用户信息失败,err:"+err.Error()), "redis查询手机号是否已经注册时候失败 RegisterOrLoginByPhoneReq：%v", in)
	}
	// TODO 调用其他服务rpc
	return &pb.UserInfoResp{
		Userinfo: &pb.UserInfoItem{
			Id:              u.Id,
			Phone:           u.Phone,
			Nickname:        u.Nickname,
			IsFollow:        false,
			Avatar:          u.Avatar,
			BackgroundImage: u.BackgroundImage,
			Signature:       u.Signature,
			TotalFavorited:  0,
			WorkCount:       0,
			FavoriteCount:   0,
			FollowCount:     0,
			FollowerCount:   0,
		},
	}, nil
}
