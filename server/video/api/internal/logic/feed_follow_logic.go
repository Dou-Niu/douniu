package logic

import (
	"context"
	"douniu/server/common/consts"
	"douniu/server/video/rpc/types/pb"
	"encoding/json"
	"github.com/pkg/errors"
	"time"

	"douniu/server/video/api/internal/svc"
	"douniu/server/video/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FeedFollowLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFeedFollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FeedFollowLogic {
	return &FeedFollowLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FeedFollowLogic) FeedFollow(req *types.FeedFollowReq) (resp *types.FeedResp, err error) {
	if req.LatestTime == 0 {
		req.LatestTime = time.Now().Unix()
	}
	meId, err := l.ctx.Value(consts.UserId).(json.Number).Int64()

	res, err := l.svcCtx.VideoRpc.FeedFollow(l.ctx, &pb.FeedFollowReq{
		LatestTime: req.LatestTime,
		MeUserId:   meId,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}
	list := res.VideoList
	videoList := make([]*types.VideoInfo, 0)
	for _, v := range list {

		videoList = append(videoList, &types.VideoInfo{
			VideoID: v.Id,
			Author: types.AuthorInfo{
				ID:              v.User.Id,
				NickName:        v.User.Nickname,
				FollowCount:     v.User.FollowCount,
				FollowerCount:   v.User.FollowerCount,
				IsFollow:        v.User.IsFollow,
				Avatar:          v.User.Avatar,
				BackgroundImage: v.User.BackgroundImage,
				Signature:       v.User.Signature,
				TotalFavorited:  v.User.TotalFavorited,
				WorkCount:       v.User.WorkCount,
				FavoriteCount:   v.User.FavoriteCount,
				CollectionCount: v.CollectionCount,
			},
			PlayURL:       v.PlayUrl,
			CoverURL:      v.CoverUrl,
			FavoriteCount: v.FavoriteCount,
			CommentCount:  v.CommentCount,
			IsFavorite:    v.IsFavorite,
			Title:         v.Title,
			Partition:     v.Partition,
			CreateTime:    time.Unix(v.CreateTime, 0).Format(time.DateTime),
		})
	}
	return &types.FeedResp{
		NextMaxValue: res.NextMaxValue,
		VideoList:    types.VideoList{List: videoList},
	}, nil

}
