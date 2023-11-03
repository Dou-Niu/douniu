package mock

import (
	"context"
	"douniu/server/relation/rpc/relationrpc"
	"google.golang.org/grpc"
)

type RelationRpc struct {
}

func (r RelationRpc) FollowAction(ctx context.Context, in *relationrpc.FollowActionRequest, opts ...grpc.CallOption) (*relationrpc.FollowActionResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (r RelationRpc) GetFollowList(ctx context.Context, in *relationrpc.GetFollowListRequest, opts ...grpc.CallOption) (*relationrpc.GetFollowListResponse, error) {
	return &relationrpc.GetFollowListResponse{}, nil
}

func (r RelationRpc) GetFollowerList(ctx context.Context, in *relationrpc.GetFollowerListRequest, opts ...grpc.CallOption) (*relationrpc.GetFollowerListResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (r RelationRpc) GetUserFollowCount(ctx context.Context, in *relationrpc.GetUserFollowCountRequest, opts ...grpc.CallOption) (*relationrpc.GetUserFollowCountResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (r RelationRpc) GetUserFollowerCount(ctx context.Context, in *relationrpc.GetUserFollowerCountRequest, opts ...grpc.CallOption) (*relationrpc.GetUserFollowerCountResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (r RelationRpc) GetFriendList(ctx context.Context, in *relationrpc.GetFriendListRequest, opts ...grpc.CallOption) (*relationrpc.GetFriendListResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (r RelationRpc) IsFollow(ctx context.Context, in *relationrpc.IsFollowRequest, opts ...grpc.CallOption) (*relationrpc.IsFollowResponse, error) {
	//TODO implement me
	panic("implement me")
}
