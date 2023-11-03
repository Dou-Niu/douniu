// Code generated by goctl. DO NOT EDIT.
// Source: favorite.proto

package server

import (
	"context"

	"douniu/server/favorite/rpc/internal/logic"
	"douniu/server/favorite/rpc/internal/svc"
	"douniu/server/favorite/rpc/pb"
)

type FavoriteRpcServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedFavoriteRpcServer
}

func NewFavoriteRpcServer(svcCtx *svc.ServiceContext) *FavoriteRpcServer {
	return &FavoriteRpcServer{
		svcCtx: svcCtx,
	}
}

func (s *FavoriteRpcServer) AddFavorite(ctx context.Context, in *pb.AddFavoriteRequest) (*pb.AddFavoriteResponse, error) {
	l := logic.NewAddFavoriteLogic(ctx, s.svcCtx)
	return l.AddFavorite(in)
}

func (s *FavoriteRpcServer) DelFavorite(ctx context.Context, in *pb.DelFavoriteRequest) (*pb.DelFavoriteResponse, error) {
	l := logic.NewDelFavoriteLogic(ctx, s.svcCtx)
	return l.DelFavorite(in)
}

func (s *FavoriteRpcServer) GetVideoFavoriteCount(ctx context.Context, in *pb.GetVideoFavoriteCountRequest) (*pb.GetVideoFavoriteCountResponse, error) {
	l := logic.NewGetVideoFavoriteCountLogic(ctx, s.svcCtx)
	return l.GetVideoFavoriteCount(in)
}

func (s *FavoriteRpcServer) GetUserFavoriteCount(ctx context.Context, in *pb.GetUserFavoriteCountRequest) (*pb.GetUserFavoriteCountResponse, error) {
	l := logic.NewGetUserFavoriteCountLogic(ctx, s.svcCtx)
	return l.GetUserFavoriteCount(in)
}

func (s *FavoriteRpcServer) GetUserFavoritedCount(ctx context.Context, in *pb.GetUserFavoritedCountRequest) (*pb.GetUserFavoritedCountResponse, error) {
	l := logic.NewGetUserFavoritedCountLogic(ctx, s.svcCtx)
	return l.GetUserFavoritedCount(in)
}

func (s *FavoriteRpcServer) IsFavorite(ctx context.Context, in *pb.IsFavoriteRequest) (*pb.IsFavoriteResponse, error) {
	l := logic.NewIsFavoriteLogic(ctx, s.svcCtx)
	return l.IsFavorite(in)
}

func (s *FavoriteRpcServer) GetFavoriteVideoIdList(ctx context.Context, in *pb.GetFavoriteVideoIdListRequest) (*pb.GetFavoriteVideoListIdResponse, error) {
	l := logic.NewGetFavoriteVideoIdListLogic(ctx, s.svcCtx)
	return l.GetFavoriteVideoIdList(in)
}

func (s *FavoriteRpcServer) AddCollection(ctx context.Context, in *pb.AddCollectionRequest) (*pb.AddCollectionResponse, error) {
	l := logic.NewAddCollectionLogic(ctx, s.svcCtx)
	return l.AddCollection(in)
}

func (s *FavoriteRpcServer) DelCollection(ctx context.Context, in *pb.DelCollectionRequest) (*pb.DelCollectionResponse, error) {
	l := logic.NewDelCollectionLogic(ctx, s.svcCtx)
	return l.DelCollection(in)
}

func (s *FavoriteRpcServer) GetUserCollectionList(ctx context.Context, in *pb.GetUserCollectionIdListRequest) (*pb.GetUserCollectionIdListResponse, error) {
	l := logic.NewGetUserCollectionListLogic(ctx, s.svcCtx)
	return l.GetUserCollectionList(in)
}

func (s *FavoriteRpcServer) GetUserCollectionIdList(ctx context.Context, in *pb.GetUserCollectionIdListRequest) (*pb.GetUserCollectionIdListResponse, error) {
	l := logic.NewGetUserCollectionIdListLogic(ctx, s.svcCtx)
	return l.GetUserCollectionIdList(in)
}

func (s *FavoriteRpcServer) GetUserCollectionCount(ctx context.Context, in *pb.GetUserCollectionCountRequest) (*pb.GetUserCollectionCountResponse, error) {
	l := logic.NewGetUserCollectionCountLogic(ctx, s.svcCtx)
	return l.GetUserCollectionCount(in)
}

func (s *FavoriteRpcServer) GetVideoCollectionCount(ctx context.Context, in *pb.GetVideoCollectionCountRequest) (*pb.GetVideoCollectionCountResponse, error) {
	l := logic.NewGetVideoCollectionCountLogic(ctx, s.svcCtx)
	return l.GetVideoCollectionCount(in)
}

func (s *FavoriteRpcServer) IsCollection(ctx context.Context, in *pb.IsCollectionRequest) (*pb.IsCollectionResponse, error) {
	l := logic.NewIsCollectionLogic(ctx, s.svcCtx)
	return l.IsCollection(in)
}
