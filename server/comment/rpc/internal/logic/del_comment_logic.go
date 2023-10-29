package logic

import (
	"context"
	"douniu/server/comment/rpc/internal/svc"
	"douniu/server/comment/rpc/pb"
	"douniu/server/common/consts"
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

	_, err = l.svcCtx.RedisClient.DecrCtx(l.ctx, consts.VideoCommentCountPrefix+strconv.Itoa(int(comment.VideoId)))
	if err != nil {
		l.Errorf("Delete comment error: %v", err)
		return
	}

	resp = new(pb.DelCommentResponse)
	return
}
