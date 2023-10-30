package mock

import (
	"context"
	"douniu/server/chat/rpc/chatrpc"
	"douniu/server/chat/rpc/pb"
	"google.golang.org/grpc"
)

type ChatRpc struct {
}

func (c ChatRpc) MessageAction(ctx context.Context, in *chatrpc.MessageActionRequest, opts ...grpc.CallOption) (*chatrpc.MessageActionResponse, error) {
	return nil, nil
}

func (c ChatRpc) MessageChat(ctx context.Context, in *chatrpc.MessageChatRequest, opts ...grpc.CallOption) (*chatrpc.MessageChatResponse, error) {
	resp := &chatrpc.MessageChatResponse{
		MessageList: []*pb.Message{
			{
				Id:         2,
				FromUserId: 1,
				ToUserId:   2,
				Content:    "123",
				CreateTime: 12345,
			}, {
				Id:         3,
				FromUserId: 2,
				ToUserId:   3,
				Content:    "456",
				CreateTime: 6789,
			},
		},
	}
	return resp, nil
}

func (c ChatRpc) MessageChatLast(ctx context.Context, in *chatrpc.MessageChatLastRequest, opts ...grpc.CallOption) (*chatrpc.MessageChatLastResponse, error) {
	resp := &chatrpc.MessageChatLastResponse{
		LastMessageList: []*pb.LastMessage{
			{
				Content: "123",
				MsgType: 456,
			},
		},
	}
	return resp, nil
}
