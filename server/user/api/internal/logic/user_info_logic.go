package logic

import (
	"context"
	"douniu/common/consts"
	"douniu/server/user/rpc/types/pb"
	"encoding/json"
	"github.com/pkg/errors"

	"douniu/server/user/api/internal/svc"
	"douniu/server/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req *types.UserInfoReq) (resp *types.UserInfoResp, err error) {
	meId, err := l.ctx.Value(consts.UserId).(json.Number).Int64()
	res, err := l.svcCtx.UserRpc.GetUserInfo(l.ctx, &pb.UserInfoReq{
		UserId:     req.UserId,
		FromUserId: meId,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}
	u := res.Userinfo
	return &types.UserInfoResp{UserInfo: &types.UserInfoItem{
		ID:              u.Id,
		Phone:           u.Phone,
		NickName:        u.Nickname,
		IsFollow:        u.IsFollow,
		Avatar:          u.Avatar,
		BackgroundImage: u.BackgroundImage,
		Signature:       u.Signature,
		TotalFavorited:  u.TotalFavorited,
		WorkCount:       u.WorkCount,
		FavoriteCount:   u.FavoriteCount,
		FollowCount:     u.FollowCount,
		FollowerCount:   u.FollowerCount,
	}}, nil
}
