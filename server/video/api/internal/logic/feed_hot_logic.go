package logic

import (
	"context"
	"douniu/common/consts"
	"douniu/common/errorx"
	"douniu/common/utils"
	"douniu/server/video/rpc/types/pb"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"math"
	"time"

	"douniu/server/video/api/internal/svc"
	"douniu/server/video/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FeedHotLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFeedHotLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FeedHotLogic {
	return &FeedHotLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FeedHotLogic) FeedHot(req *types.FeedHotReq) (resp *types.FeedHotResp, err error) {
	err = utils.DefaultGetValidParams(l.ctx, req)
	if err != nil {
		return nil, errors.Wrapf(errorx.NewCodeError(1, fmt.Sprintf("validate校验错误: %v", err)), "validate校验错误err :%v", err)
	}
	if req.MaxHot == 0 {
		req.MaxHot = math.MaxInt
	}
	meId, err := l.ctx.Value(consts.UserId).(json.Number).Int64()

	res, err := l.svcCtx.VideoRpc.FeedHot(l.ctx, &pb.FeedHotReq{
		MaxHot:   req.MaxHot,
		MeUserId: meId,
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
	return &types.FeedHotResp{
		NextMaxHot: res.NextMaxHost,
		VideoList:  types.VideoList{List: videoList},
	}, nil

}
