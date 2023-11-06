package logic

import (
	"context"
	"douniu/common/errorx"
	"fmt"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logc"

	"douniu/server/video/rpc/internal/svc"
	"douniu/server/video/rpc/types/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAuthorIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAuthorIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAuthorIdLogic {
	return &GetAuthorIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAuthorIdLogic) GetAuthorId(in *pb.GetAuthorIdReq) (*pb.GetAuthorIdResp, error) {
	video, err := l.svcCtx.VideoModel.FindOne(l.ctx, in.VideoId)
	if err != nil {
		logc.Error(l.ctx, fmt.Sprintf("根据video_id查询user_id错误,err:%v", err))
		return nil, errors.Wrapf(errorx.NewDefaultError("根据video_id查询user_id错误"), "注册失败，userinfo 写入mysql错误 RegisterReq：%v", in)
	}
	return &pb.GetAuthorIdResp{
		UserId: video.UserId,
	}, nil
}
