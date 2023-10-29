package logic

import (
	"context"
	"douniu/server/comment/api/internal/svc"
	"douniu/server/comment/api/internal/types"
	"douniu/server/comment/rpc/commentrpc"
	"douniu/server/common/consts"
	"encoding/json"
	"github.com/jinzhu/copier"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommentDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommentDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentDetailLogic {
	return &CommentDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommentDetailLogic) CommentDetail(req *types.CommentDetailRequest) (resp *types.CommentDetailResponse, err error) {
	userId, _ := l.ctx.Value(consts.UserId).(json.Number).Int64()
	resp = new(types.CommentDetailResponse)
	detailResp, err := l.svcCtx.CommentRpc.GetCommentDetail(l.ctx, &commentrpc.GetCommentDetailRequest{CommentId: req.CommentId, UserId: userId})
	if err != nil {
		return
	}
	resp.CommentList = make([]*types.Comment, 0, len(detailResp.CommentList))
	_ = copier.Copy(&resp.CommentList, detailResp.CommentList)

	return
}
