package logic

import (
	"context"
	"douniu/server/common/errorx"
	"douniu/server/favorite/rpc/favoriterpc"
	"douniu/server/relation/rpc/relationrpc"
	"douniu/server/video/rpc/videorpc"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/mr"

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
	var favoriteCount, totalFavorited, collectionCount, workCount, followCount, followerCount int64
	var isFollow bool
	err = mr.Finish(func() error {
		// 获取用户总的点赞数
		res, err := l.svcCtx.FavoriteRpc.GetUserFavoriteCount(l.ctx, &favoriterpc.GetUserFavoriteCountRequest{
			UserId: in.UserId,
		})
		favoriteCount = res.Count
		return err
	}, func() error {
		// 获取用户被点赞数
		res, err := l.svcCtx.FavoriteRpc.GetUserFavoritedCount(l.ctx, &favoriterpc.GetUserFavoritedCountRequest{
			UserId: in.UserId,
		})
		totalFavorited = res.Count
		return err
	}, func() error {
		// 用户用户总的收藏数
		res, err := l.svcCtx.FavoriteRpc.GetUserCollectionCount(l.ctx, &favoriterpc.GetUserCollectionCountRequest{
			UserId: in.UserId,
		})
		collectionCount = res.Count
		return err
	}, func() error {
		// 获取是否关注该用户
		res, err := l.svcCtx.RelationRpc.IsFollow(l.ctx, &relationrpc.IsFollowRequest{
			UserId:       in.FromUserId,
			TargetUserId: in.UserId,
		})
		isFollow = res.IsFollow
		return err
	}, func() error {
		// 获取用户视频总数
		res, err := l.svcCtx.VideoRpc.GetWorkCount(l.ctx, &videorpc.WorkCountReq{
			UserId: in.UserId,
		})
		workCount = res.WorkCount
		return err
	}, func() error {
		// 获取关注总数
		res, err := l.svcCtx.RelationRpc.GetUserFollowCount(l.ctx, &relationrpc.GetUserFollowCountRequest{
			UserId: in.UserId,
		})
		followCount = res.Count
		return err
	}, func() error {
		// 获取粉丝数
		res, err := l.svcCtx.RelationRpc.GetUserFollowerCount(l.ctx, &relationrpc.GetUserFollowerCountRequest{
			UserId: in.UserId,
		})
		followerCount = res.Count
		return err
	})
	if err != nil {
		// Handle the error, log, and return if needed
		logc.Error(l.ctx, err, "RPC call error")
		return nil, errors.Wrapf(err, "req: %+v", in)

	}
	return &pb.UserInfoResp{
		Userinfo: &pb.UserInfoItem{
			Id:              u.Id,
			Phone:           u.Phone,
			Nickname:        u.Nickname,
			IsFollow:        isFollow,
			Avatar:          u.Avatar,
			BackgroundImage: u.BackgroundImage,
			Signature:       u.Signature,
			TotalFavorited:  totalFavorited,
			CollectionCount: collectionCount,
			WorkCount:       workCount,
			FavoriteCount:   favoriteCount,
			FollowCount:     followCount,
			FollowerCount:   followerCount,
		},
	}, nil
}
