package logic

import (
	"context"
	consts2 "douniu/common/consts"
	"douniu/server/chat/model"
	"google.golang.org/protobuf/proto"
	"strconv"
	"time"

	"douniu/server/chat/rpc/internal/svc"
	"douniu/server/chat/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type MessageActionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMessageActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MessageActionLogic {
	return &MessageActionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MessageActionLogic) MessageAction(in *pb.MessageActionRequest) (resp *pb.MessageActionResponse, err error) {
	message := &model.Message{
		Id:         l.svcCtx.Snowflake.Generate().Int64(),
		FromUserId: in.FromUserId,
		ToUserId:   in.ToUserId,
		Content:    in.Content,
		CreateTime: time.Now().Unix(),
	}
	//messageStr, err := jsonx.MarshalToString(message)
	////err = l.svcCtx.KafkaPusher.Push(messageStr)
	//if err != nil {
	//	return nil, err
	//}

	_, err = l.svcCtx.MessageModel.Insert(context.Background(), message)
	if err != nil {
		logx.Errorf("MessageAction MessageModel.Insert error: %s", err.Error())
		return
	}

	// 保存最新消息到redis
	fromUserID := strconv.Itoa(int(message.FromUserId))
	toUserID := strconv.Itoa(int(message.ToUserId))

	lastMessage := &pb.LastMessage{Content: message.Content}
	lastMessage.MsgType = consts2.MsgTypeRecv
	lastMessageRecvBytes, _ := proto.Marshal(lastMessage)
	//lastMessageRecvStr, _ := jsonx.MarshalToString(lastMessage)
	lastMessage.MsgType = consts2.MsgTypeSend
	lastMessageSendBytes, _ := proto.Marshal(lastMessage)
	//lastMessageSendStr, err := jsonx.MarshalToString(lastMessage)

	_ = l.svcCtx.RedisClient.Hset(consts2.LastMessagePrefix+fromUserID, toUserID, string(lastMessageSendBytes))
	_ = l.svcCtx.RedisClient.Hset(consts2.LastMessagePrefix+toUserID, fromUserID, string(lastMessageRecvBytes))

	resp = new(pb.MessageActionResponse)

	return
}
