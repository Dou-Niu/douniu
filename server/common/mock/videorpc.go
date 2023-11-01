package mock

import (
	"context"
	"douniu/server/video/rpc/videorpc"
	"google.golang.org/grpc"
)

type VideoRpc struct {
}

func (v VideoRpc) PublishVideo(ctx context.Context, in *videorpc.PublishVideoReq, opts ...grpc.CallOption) (*videorpc.CommonResp, error) {
	//TODO implement me
	panic("implement me")
}

func (v VideoRpc) FeedHome(ctx context.Context, in *videorpc.FeedHomeReq, opts ...grpc.CallOption) (*videorpc.FeedHomeResp, error) {
	//TODO implement me
	panic("implement me")
}

func (v VideoRpc) FeedHot(ctx context.Context, in *videorpc.FeedHotReq, opts ...grpc.CallOption) (*videorpc.FeedHotResp, error) {
	//TODO implement me
	panic("implement me")
}

func (v VideoRpc) FeedUser(ctx context.Context, in *videorpc.FeedUserReq, opts ...grpc.CallOption) (*videorpc.FeedResp, error) {
	//TODO implement me
	panic("implement me")
}

func (v VideoRpc) FeedPartition(ctx context.Context, in *videorpc.FeedPartitionReq, opts ...grpc.CallOption) (*videorpc.FeedResp, error) {
	//TODO implement me
	panic("implement me")
}

func (v VideoRpc) SearchVideo(ctx context.Context, in *videorpc.SearchVideoReq, opts ...grpc.CallOption) (*videorpc.FeedResp, error) {
	//TODO implement me
	panic("implement me")
}

func (v VideoRpc) DeleteVideo(ctx context.Context, in *videorpc.DeleteVideoReq, opts ...grpc.CallOption) (*videorpc.CommonResp, error) {
	//TODO implement me
	panic("implement me")
}

func (v VideoRpc) GetVideoListInfo(ctx context.Context, in *videorpc.GetVideoListInfoReq, opts ...grpc.CallOption) (*videorpc.GetVideoListInfoResp, error) {
	return &videorpc.GetVideoListInfoResp{
		VideoList: []*videorpc.Video{
			{
				Id: 1,
				User: &videorpc.User{
					Id:              2,
					Nickname:        "张三",
					FollowCount:     2,
					FollowerCount:   200,
					IsFollow:        true,
					Avatar:          "httpaaa",
					BackgroundImage: "http123",
					Signature:       "aaa",
					TotalFavorited:  1,
					WorkCount:       2,
					FavoriteCount:   3,
				},
				PlayUrl:       "httpppp",
				CoverUrl:      "htppqweq",
				FavoriteCount: 2,
				CommentCount:  3,
				IsFavorite:    true,
				Title:         "标题",
			}, {
				Id: 2,
				User: &videorpc.User{
					Id:              2,
					Nickname:        "张三",
					FollowCount:     2,
					FollowerCount:   200,
					IsFollow:        true,
					Avatar:          "httpaaa",
					BackgroundImage: "http123",
					Signature:       "aaa",
					TotalFavorited:  1,
					WorkCount:       2,
					FavoriteCount:   3,
				},
				PlayUrl:       "httpppp",
				CoverUrl:      "htppqweq",
				FavoriteCount: 2,
				CommentCount:  3,
				IsFavorite:    true,
				Title:         "标题",
			},
		},
	}, nil
}

func (v VideoRpc) GetAuthorId(ctx context.Context, in *videorpc.GetAuthorIdReq, opts ...grpc.CallOption) (*videorpc.GetAuthorIdResp, error) {
	return &videorpc.GetAuthorIdResp{UserId: 3}, nil
}

func NewVideoRpc() VideoRpc {
	return VideoRpc{}
}
