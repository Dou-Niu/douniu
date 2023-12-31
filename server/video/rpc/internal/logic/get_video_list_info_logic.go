package logic

import (
	"context"
	"douniu/server/comment/rpc/commentrpc"
	"douniu/server/favorite/rpc/favoriterpc"
	"douniu/server/user/rpc/userrpc"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/mr"

	"douniu/server/video/rpc/internal/svc"
	"douniu/server/video/rpc/types/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetVideoListInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetVideoListInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetVideoListInfoLogic {
	return &GetVideoListInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetVideoListInfoLogic) GetVideoListInfo(in *pb.GetVideoListInfoReq) (*pb.GetVideoListInfoResp, error) {
	// 取出视频信息
	res := make([]*pb.Video, 0)
	for _, v := range in.VideoIdList {
		oneVideo, err := l.svcCtx.VideoModel.FindOne(l.ctx, v)
		if err != nil {
			logc.Error(l.ctx, err)
			// 这里不直接return 防止因为某一个id没查到导致全部卡死
			continue
			//return nil, errors.Wrapf(errorx.NewDefaultError("mysql查询视频失败"+err.Error()), "更改用户密码失败ResetPassword：%v", in)
		}
		var userInfo *userrpc.UserInfoItem
		var favoriteCount, collectionCount, commentCount int64
		var isFavorite, isCollect bool
		err = mr.Finish(func() error {
			// 获取视频作者信息
			res, err := l.svcCtx.UserRpc.GetUserInfo(l.ctx, &userrpc.UserInfoReq{
				UserId:     oneVideo.UserId,
				FromUserId: in.MeUserId,
			})
			if err != nil {
				return errors.New("获取视频作者信息错误" + err.Error())
			}
			userInfo = res.Userinfo
			return nil
		}, func() error {
			// 获取视频获赞数
			res, err := l.svcCtx.FavoriteRpc.GetVideoFavoriteCount(l.ctx, &favoriterpc.GetVideoFavoriteCountRequest{
				VideoId: oneVideo.Id,
			})
			if err != nil {
				return errors.New("获取视频获赞数错误" + err.Error())
			}
			favoriteCount = res.Count
			return nil
		}, func() error {
			// 获取视频收藏数
			res, err := l.svcCtx.FavoriteRpc.GetVideoCollectionCount(l.ctx, &favoriterpc.GetVideoCollectionCountRequest{
				VideoId: oneVideo.Id,
			})
			if err != nil {
				return errors.New("获取视频收藏数错误" + err.Error())
			}
			collectionCount = res.Count

			return nil
		}, func() error {
			// 是否对视频点赞
			res, err := l.svcCtx.FavoriteRpc.IsFavorite(l.ctx, &favoriterpc.IsFavoriteRequest{
				UserId:  in.MeUserId,
				VideoId: oneVideo.Id,
			})
			if err != nil {
				return errors.New("获取是否对视频点赞错误" + err.Error())
			}
			isFavorite = res.IsFavorite
			return nil
		}, func() error {
			// 是否对视频收藏
			res, err := l.svcCtx.FavoriteRpc.IsCollection(l.ctx, &favoriterpc.IsCollectionRequest{
				UserId:  in.MeUserId,
				VideoId: oneVideo.Id,
			})
			if err != nil {
				return errors.New("获取是否对视频收藏错误" + err.Error())
			}
			isCollect = res.IsCollection
			return nil
		}, func() error {
			// 获取视频的评论数
			res, err := l.svcCtx.CommentRpc.GetCommentCount(l.ctx, &commentrpc.GetCommentCountRequest{
				VideoId: oneVideo.Id,
			})
			if err != nil {
				return errors.New("获取视频的评论数错误" + err.Error())
			}
			commentCount = res.Count
			return nil
		})
		if err != nil {
			// Handle the error, log, and return if needed
			logc.Error(l.ctx, err, "RPC call error")
			return nil, errors.Wrapf(err, "req: %+v", in)

		}
		res = append(res, &pb.Video{
			Id: oneVideo.Id,
			User: &pb.User{
				Id:              userInfo.Id,
				Nickname:        userInfo.Nickname,
				FollowCount:     userInfo.FollowCount,
				FollowerCount:   userInfo.FollowerCount,
				IsFollow:        userInfo.IsFollow,
				Avatar:          userInfo.Avatar,
				BackgroundImage: userInfo.BackgroundImage,
				Signature:       userInfo.Signature,
				TotalFavorited:  userInfo.TotalFavorited,
				CollectionCount: userInfo.CollectionCount,
				WorkCount:       userInfo.WorkCount,
				FavoriteCount:   userInfo.FavoriteCount,
			},
			PlayUrl:         oneVideo.PlayUrl,
			CoverUrl:        oneVideo.CoverUrl,
			FavoriteCount:   favoriteCount,
			CollectionCount: collectionCount,
			CommentCount:    commentCount,
			IsFavorite:      isFavorite,
			IsCollect:       isCollect,
			Title:           oneVideo.Title,
			Partition:       oneVideo.Partition,
			CreateTime:      oneVideo.CreateTime.Unix(),
		})
	}

	return &pb.GetVideoListInfoResp{
		VideoList: res,
	}, nil
}
