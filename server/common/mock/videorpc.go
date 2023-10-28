package mock

import (
	"context"
	"douniu/server/video/videorpc"
	"google.golang.org/grpc"
)

type VideoRpc struct {
}

func (v VideoRpc) GetVideoListInfo(ctx context.Context, in *videorpc.GetVideoListInfoReq, opts ...grpc.CallOption) (*videorpc.GetVideoListInfoResp, error) {
	return &videorpc.GetVideoListInfoResp{
		VideoList: []*videorpc.Video{
			{
				Id: 1,
				User: &videorpc.User{
					Id:              2,
					Username:        "张三",
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
					Username:        "张三",
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
