package logic

import (
	"context"
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
			// 这里不直接return 防止因为某一个id没查到导致全部卡死
			continue
			//return nil, errors.Wrapf(errorx.NewDefaultError("mysql查询视频失败"+err.Error()), "更改用户密码失败ResetPassword：%v", in)
			// TODO 调用其他rpc
		}
		res = append(res, &pb.Video{
			Id:              oneVideo.Id,
			User:            nil,
			PlayUrl:         oneVideo.PlayUrl,
			CoverUrl:        oneVideo.CoverUrl,
			FavoriteCount:   0,
			CollectionCount: 0,
			CommentCount:    0,
			IsFavorite:      false,
			Title:           oneVideo.Title,
			Partition:       oneVideo.Partition,
			CreateTime:      oneVideo.CreateTime.Unix(),
		})
	}

	return &pb.GetVideoListInfoResp{
		VideoList: res,
	}, nil
}
