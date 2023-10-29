package logic

import (
	"context"
	"douniu/server/comment/rpc/commentrpc"
	"douniu/server/common/consts"
	"encoding/json"
	"github.com/jinzhu/copier"

	"douniu/server/comment/api/internal/svc"
	"douniu/server/comment/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommentListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentListLogic {
	return &CommentListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommentListLogic) CommentList(req *types.CommentListRequest) (resp *types.CommentListResponse, err error) {
	userId, _ := l.ctx.Value(consts.UserId).(json.Number).Int64()
	resp = new(types.CommentListResponse)

	commentListResp, err := l.svcCtx.CommentRpc.GetCommentList(l.ctx, &commentrpc.GetCommentListRequest{
		VideoId: req.VideoId,
		UserId:  userId,
	})

	if err != nil {
		l.Errorf("Get comment list error: %v", err)
		return nil, err
	}
	resp.CommentList = make([]*types.Comment, len(commentListResp.CommentList))
	_ = copier.Copy(&resp.CommentList, &commentListResp.CommentList)

	return
}
