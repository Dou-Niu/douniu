package logic

import (
	"context"
	"douniu/server/common/errorx"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logc"

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
			return nil, errors.Wrapf(errorx.NewDefaultError("mysql查询视频失败"+err.Error()), "更改用户密码失败ResetPassword：%v", in)
			// TODO 调用其他rpc
		}
		res = append(res, &pb.Video{
			Id:            oneVideo.Id,
			User:          nil,
			PlayUrl:       oneVideo.PlayUrl,
			CoverUrl:      oneVideo.CoverUrl,
			FavoriteCount: 0,
			CommentCount:  0,
			IsFavorite:    false,
			Title:         oneVideo.Title,
			Partition:     oneVideo.Partition,
		})
	}

	return &pb.GetVideoListInfoResp{
		VideoList: res,
	}, nil
}
