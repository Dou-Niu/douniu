package logic

import (
	"context"
	"douniu/common/consts"
	"douniu/server/comment/rpc/commentrpc"
	"encoding/json"
	"errors"
	"github.com/jinzhu/copier"

	"douniu/server/comment/api/internal/svc"
	"douniu/server/comment/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommentActionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommentActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentActionLogic {
	return &CommentActionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommentActionLogic) CommentAction(req *types.CommentActionRequest) (resp *types.CommentActionResponse, err error) {
	userId, _ := l.ctx.Value(consts.UserId).(json.Number).Int64()
	resp = new(types.CommentActionResponse)
	if req.ActionType == consts.CommentAdd {
		addCommentResp, err := l.svcCtx.CommentRpc.AddComment(l.ctx, &commentrpc.AddCommentRequest{
			UserId:   userId,
			VideoId:  req.VideoId,
			Content:  req.Content,
			ParentId: req.ParentId,
		})
		if err != nil {
			return nil, err
		}
		resp.Comment = new(types.Comment)
		_ = copier.Copy(&resp.Comment, &addCommentResp.Comment)

	} else if req.ActionType == consts.CommentDel {
		if req.CommentId <= 0 {
			return nil, errors.New("id不合法")
		}

		_, err := l.svcCtx.CommentRpc.DelComment(l.ctx, &commentrpc.DelCommentRequest{
			CommentId: req.CommentId,
			VideoId:   req.VideoId,
		})
		if err != nil {
			return nil, err
		}

	}

	return
}
