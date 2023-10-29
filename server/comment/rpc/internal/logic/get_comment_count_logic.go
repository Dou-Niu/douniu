package logic

import (
	"context"
	"douniu/server/comment/rpc/internal/svc"
	"douniu/server/comment/rpc/pb"
	"douniu/server/common/consts"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommentCountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCommentCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentCountLogic {
	return &GetCommentCountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCommentCountLogic) GetCommentCount(in *pb.GetCommentCountRequest) (resp *pb.GetCommentCountResponse, err error) {
	resp = new(pb.GetCommentCountResponse)
	var count int
	countStr, err := l.svcCtx.RedisClient.GetCtx(l.ctx, consts.VideoCommentCountPrefix+strconv.Itoa(int(in.VideoId)))
	if err != nil {
		l.Errorf("Get comment count error: %v", err)
		// redis出错，从数据库中获取数量
		resp.Count, err = l.svcCtx.CommentModel.GetCountByVideoId(l.ctx, in.VideoId)
		if err != nil {
			l.Errorf("Get comment count error: %v", err)
			return
		}

	} else {
		if countStr == "" {
			countStr = "0"
		}
		count, err = strconv.Atoi(countStr)
		if err != nil {
			l.Errorf("Get comment count error: %v", err)
			return
		}

		resp.Count = int64(count)
	}

	return
}
