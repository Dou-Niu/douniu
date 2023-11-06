package logic

import (
	"context"
	"douniu/common/consts"
	"douniu/common/errorx"
	"douniu/server/video/rpc/types/pb"
	"encoding/json"
	"github.com/pkg/errors"
	"time"

	"douniu/server/video/api/internal/svc"
	"douniu/server/video/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LookOneVideoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLookOneVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LookOneVideoLogic {
	return &LookOneVideoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LookOneVideoLogic) LookOneVideo(req *types.LookOneVideoReq) (resp *types.FeedResp, err error) {
	meId, err := l.ctx.Value(consts.UserId).(json.Number).Int64()

	res, err := l.svcCtx.VideoRpc.GetVideoListInfo(l.ctx, &pb.GetVideoListInfoReq{
		VideoIdList: []int64{req.VideoId},
		MeUserId:    meId,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}
	list := res.VideoList
	if len(list) == 0 {
		return nil, errors.Wrapf(errorx.NewDefaultError("查不到数据了"), "req: %+v", req)

	}
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
		IsFinal:   true,
		VideoList: types.VideoList{List: videoList},
	}, nil
}
