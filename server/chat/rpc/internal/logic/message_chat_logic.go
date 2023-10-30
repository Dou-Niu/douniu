package logic

import (
	"context"

	"douniu/server/chat/rpc/internal/svc"
	"douniu/server/chat/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type MessageChatLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMessageChatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MessageChatLogic {
	return &MessageChatLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MessageChatLogic) MessageChat(in *pb.MessageChatRequest) (resp *pb.MessageChatResponse, err error) {
	// todo: add your logic here and delete this line

	resp = new(pb.MessageChatResponse)

	return
}
