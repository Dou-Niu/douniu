package logic

import (
	"context"
	consts2 "douniu/common/consts"
	"douniu/server/comment/rpc/internal/svc"
	"douniu/server/comment/rpc/pb"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelCommentLogic {
	return &DelCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelCommentLogic) DelComment(in *pb.DelCommentRequest) (resp *pb.DelCommentResponse, err error) {
	comment, err := l.svcCtx.CommentModel.FindOne(l.ctx, in.CommentId)
	if errors.Is(err, sqlx.ErrNotFound) {
		return nil, errors.New("评论不存在")
	}

	err = l.svcCtx.CommentModel.Delete(l.ctx, in.CommentId)
	if err != nil && !errors.Is(err, sqlx.ErrNotFound) {
		l.Errorf("Delete comment error: %v", err)
		return
	}

	_, err = l.svcCtx.RedisClient.DecrCtx(l.ctx, consts2.VideoCommentCountPrefix+strconv.Itoa(int(comment.VideoId)))
	if err != nil {
		l.Errorf("Delete comment error: %v", err)
		return
	}

	// 视频热度减少
	videoIdStr := strconv.Itoa(int(in.VideoId))
	_, err = l.svcCtx.RedisClient.ZincrbyCtx(l.ctx, consts2.VideoHotScore, -int64(consts2.SingleHotScore), videoIdStr)
	if err != nil {
		l.Errorf("RedisClient ZincrbyCtx error: %v", err)
		return
	}

	resp = new(pb.DelCommentResponse)
	return
}
