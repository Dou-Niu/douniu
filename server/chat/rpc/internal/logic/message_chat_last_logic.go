package logic

import (
	"context"

	"douniu/server/chat/rpc/internal/svc"
	"douniu/server/chat/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type MessageChatLastLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMessageChatLastLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MessageChatLastLogic {
	return &MessageChatLastLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MessageChatLastLogic) MessageChatLast(in *pb.MessageChatLastRequest) (resp *pb.MessageChatLastResponse, err error) {
	// todo: add your logic here and delete this line

	resp = new(pb.MessageChatLastResponse)

	return
}
