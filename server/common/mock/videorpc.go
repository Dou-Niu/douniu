package mock

import (
	"context"
	"douniu/server/video/videorpc"
	"google.golang.org/grpc"
)

type VideoRpc struct {
}

func (v VideoRpc) GetAuthorId(ctx context.Context, in *videorpc.GetAuthorIdReq, opts ...grpc.CallOption) (*videorpc.GetAuthorIdResp, error) {
	return &videorpc.GetAuthorIdResp{UserId: 3}, nil
}

func NewVideoRpc() VideoRpc {
	return VideoRpc{}
}
