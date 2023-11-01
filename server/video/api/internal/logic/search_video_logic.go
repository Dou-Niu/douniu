package logic

import (
	"context"
	"douniu/server/common/consts"
	"douniu/server/common/errorx"
	"douniu/server/common/utils"
	"douniu/server/video/rpc/types/pb"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"time"

	"douniu/server/video/api/internal/svc"
	"douniu/server/video/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchVideoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchVideoLogic {
	return &SearchVideoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchVideoLogic) SearchVideo(req *types.SearchVideoReq) (resp *types.FeedResp, err error) {
	err = utils.DefaultGetValidParams(l.ctx, req)
	if err != nil {
		return nil, errors.Wrapf(errorx.NewCodeError(1, fmt.Sprintf("validate校验错误: %v", err)), "validate校验错误err :%v", err)
	}
	meId, err := l.ctx.Value(consts.UserId).(json.Number).Int64()
	res, err := l.svcCtx.VideoRpc.SearchVideo(l.ctx, &pb.SearchVideoReq{
		MeUserId: meId,
		KeyWords: req.KeyWords,
		Page:     req.Page,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}
	list := res.VideoList
	videoList := make([]*types.VideoInfo, 0)
	for _, v := range list {

		videoList = append(videoList, &types.VideoInfo{
			VideoID: v.Id,
			Author:  types.AuthorInfo{
				//ID:              v.User.Id,
				//NickName:        v.User.Nickname,
				//FollowCount:     v.User.FollowCount,
				//FollowerCount:   v.User.FollowerCount,
				//IsFollow:        v.User.IsFollow,
				//Avatar:          v.User.Avatar,
				//BackgroundImage: v.User.BackgroundImage,
				//Signature:       v.User.Signature,
				//TotalFavorited:  v.User.TotalFavorited,
				//WorkCount:       v.User.WorkCount,
				//FavoriteCount:   v.User.FavoriteCount,

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
		IsFinal:      res.IsFinal,
	}, nil
}
