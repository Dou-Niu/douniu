package logic

import (
	"context"
	"douniu/server/chat/rpc/pb"
	"douniu/server/common/consts"
	"encoding/json"
	"errors"

	"douniu/server/chat/api/internal/svc"
	"douniu/server/chat/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MessageActionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMessageActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MessageActionLogic {
	return &MessageActionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MessageActionLogic) MessageAction(req *types.MessageActionRequest) (resp *types.MessageActionResponse, err error) {
	fromUserId, _ := l.ctx.Value(consts.UserId).(json.Number).Int64()

	if fromUserId == req.ToUserId {
		return nil, errors.New("不能给自己发消息")
	}

	_, err = l.svcCtx.ChatRpc.MessageAction(l.ctx, &pb.MessageActionRequest{
		FromUserId: fromUserId,
		ToUserId:   req.ToUserId,
		Action:     req.ActionType,
		Content:    req.Content,
	})
	if err != nil {
		return
	}

	resp = new(types.MessageActionResponse)

	return
}
