// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.10
// source: favorite.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	FavoriteRpc_AddFavorite_FullMethodName             = "/favorite.FavoriteRpc/AddFavorite"
	FavoriteRpc_DelFavorite_FullMethodName             = "/favorite.FavoriteRpc/DelFavorite"
	FavoriteRpc_GetVideoFavoriteCount_FullMethodName   = "/favorite.FavoriteRpc/GetVideoFavoriteCount"
	FavoriteRpc_GetUserFavoriteCount_FullMethodName    = "/favorite.FavoriteRpc/GetUserFavoriteCount"
	FavoriteRpc_GetUserFavoritedCount_FullMethodName   = "/favorite.FavoriteRpc/GetUserFavoritedCount"
	FavoriteRpc_IsFavorite_FullMethodName              = "/favorite.FavoriteRpc/IsFavorite"
	FavoriteRpc_GetFavoriteVideoIdList_FullMethodName  = "/favorite.FavoriteRpc/GetFavoriteVideoIdList"
	FavoriteRpc_AddCollection_FullMethodName           = "/favorite.FavoriteRpc/AddCollection"
	FavoriteRpc_DelCollection_FullMethodName           = "/favorite.FavoriteRpc/DelCollection"
	FavoriteRpc_GetUserCollectionList_FullMethodName   = "/favorite.FavoriteRpc/GetUserCollectionList"
	FavoriteRpc_GetUserCollectionCount_FullMethodName  = "/favorite.FavoriteRpc/GetUserCollectionCount"
	FavoriteRpc_GetVideoCollectionCount_FullMethodName = "/favorite.FavoriteRpc/GetVideoCollectionCount"
	FavoriteRpc_IsCollection_FullMethodName            = "/favorite.FavoriteRpc/IsCollection"
)

// FavoriteRpcClient is the client API for FavoriteRpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FavoriteRpcClient interface {
	AddFavorite(ctx context.Context, in *AddFavoriteRequest, opts ...grpc.CallOption) (*AddFavoriteResponse, error)
	DelFavorite(ctx context.Context, in *DelFavoriteRequest, opts ...grpc.CallOption) (*DelFavoriteResponse, error)
	GetVideoFavoriteCount(ctx context.Context, in *GetVideoFavoriteCountRequest, opts ...grpc.CallOption) (*GetVideoFavoriteCountResponse, error)
	GetUserFavoriteCount(ctx context.Context, in *GetUserFavoriteCountRequest, opts ...grpc.CallOption) (*GetUserFavoriteCountResponse, error)
	GetUserFavoritedCount(ctx context.Context, in *GetUserFavoritedCountRequest, opts ...grpc.CallOption) (*GetUserFavoritedCountResponse, error)
	IsFavorite(ctx context.Context, in *IsFavoriteRequest, opts ...grpc.CallOption) (*IsFavoriteResponse, error)
	GetFavoriteVideoIdList(ctx context.Context, in *GetFavoriteVideoIdListRequest, opts ...grpc.CallOption) (*GetFavoriteVideoListIdResponse, error)
	AddCollection(ctx context.Context, in *AddCollectionRequest, opts ...grpc.CallOption) (*AddCollectionResponse, error)
	DelCollection(ctx context.Context, in *DelCollectionRequest, opts ...grpc.CallOption) (*DelCollectionResponse, error)
	GetUserCollectionList(ctx context.Context, in *GetUserCollectionIdListRequest, opts ...grpc.CallOption) (*GetUserCollectionIdListResponse, error)
	GetUserCollectionCount(ctx context.Context, in *GetUserCollectionCountRequest, opts ...grpc.CallOption) (*GetUserCollectionCountResponse, error)
	GetVideoCollectionCount(ctx context.Context, in *GetVideoCollectionCountRequest, opts ...grpc.CallOption) (*GetVideoCollectionCountResponse, error)
	IsCollection(ctx context.Context, in *IsCollectionRequest, opts ...grpc.CallOption) (*IsCollectionResponse, error)
}

type favoriteRpcClient struct {
	cc grpc.ClientConnInterface
}

func NewFavoriteRpcClient(cc grpc.ClientConnInterface) FavoriteRpcClient {
	return &favoriteRpcClient{cc}
}

func (c *favoriteRpcClient) AddFavorite(ctx context.Context, in *AddFavoriteRequest, opts ...grpc.CallOption) (*AddFavoriteResponse, error) {
	out := new(AddFavoriteResponse)
	err := c.cc.Invoke(ctx, FavoriteRpc_AddFavorite_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *favoriteRpcClient) DelFavorite(ctx context.Context, in *DelFavoriteRequest, opts ...grpc.CallOption) (*DelFavoriteResponse, error) {
	out := new(DelFavoriteResponse)
	err := c.cc.Invoke(ctx, FavoriteRpc_DelFavorite_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *favoriteRpcClient) GetVideoFavoriteCount(ctx context.Context, in *GetVideoFavoriteCountRequest, opts ...grpc.CallOption) (*GetVideoFavoriteCountResponse, error) {
	out := new(GetVideoFavoriteCountResponse)
	err := c.cc.Invoke(ctx, FavoriteRpc_GetVideoFavoriteCount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *favoriteRpcClient) GetUserFavoriteCount(ctx context.Context, in *GetUserFavoriteCountRequest, opts ...grpc.CallOption) (*GetUserFavoriteCountResponse, error) {
	out := new(GetUserFavoriteCountResponse)
	err := c.cc.Invoke(ctx, FavoriteRpc_GetUserFavoriteCount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *favoriteRpcClient) GetUserFavoritedCount(ctx context.Context, in *GetUserFavoritedCountRequest, opts ...grpc.CallOption) (*GetUserFavoritedCountResponse, error) {
	out := new(GetUserFavoritedCountResponse)
	err := c.cc.Invoke(ctx, FavoriteRpc_GetUserFavoritedCount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *favoriteRpcClient) IsFavorite(ctx context.Context, in *IsFavoriteRequest, opts ...grpc.CallOption) (*IsFavoriteResponse, error) {
	out := new(IsFavoriteResponse)
	err := c.cc.Invoke(ctx, FavoriteRpc_IsFavorite_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *favoriteRpcClient) GetFavoriteVideoIdList(ctx context.Context, in *GetFavoriteVideoIdListRequest, opts ...grpc.CallOption) (*GetFavoriteVideoListIdResponse, error) {
	out := new(GetFavoriteVideoListIdResponse)
	err := c.cc.Invoke(ctx, FavoriteRpc_GetFavoriteVideoIdList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *favoriteRpcClient) AddCollection(ctx context.Context, in *AddCollectionRequest, opts ...grpc.CallOption) (*AddCollectionResponse, error) {
	out := new(AddCollectionResponse)
	err := c.cc.Invoke(ctx, FavoriteRpc_AddCollection_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *favoriteRpcClient) DelCollection(ctx context.Context, in *DelCollectionRequest, opts ...grpc.CallOption) (*DelCollectionResponse, error) {
	out := new(DelCollectionResponse)
	err := c.cc.Invoke(ctx, FavoriteRpc_DelCollection_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *favoriteRpcClient) GetUserCollectionList(ctx context.Context, in *GetUserCollectionIdListRequest, opts ...grpc.CallOption) (*GetUserCollectionIdListResponse, error) {
	out := new(GetUserCollectionIdListResponse)
	err := c.cc.Invoke(ctx, FavoriteRpc_GetUserCollectionList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *favoriteRpcClient) GetUserCollectionCount(ctx context.Context, in *GetUserCollectionCountRequest, opts ...grpc.CallOption) (*GetUserCollectionCountResponse, error) {
	out := new(GetUserCollectionCountResponse)
	err := c.cc.Invoke(ctx, FavoriteRpc_GetUserCollectionCount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *favoriteRpcClient) GetVideoCollectionCount(ctx context.Context, in *GetVideoCollectionCountRequest, opts ...grpc.CallOption) (*GetVideoCollectionCountResponse, error) {
	out := new(GetVideoCollectionCountResponse)
	err := c.cc.Invoke(ctx, FavoriteRpc_GetVideoCollectionCount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *favoriteRpcClient) IsCollection(ctx context.Context, in *IsCollectionRequest, opts ...grpc.CallOption) (*IsCollectionResponse, error) {
	out := new(IsCollectionResponse)
	err := c.cc.Invoke(ctx, FavoriteRpc_IsCollection_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FavoriteRpcServer is the server API for FavoriteRpc service.
// All implementations must embed UnimplementedFavoriteRpcServer
// for forward compatibility
type FavoriteRpcServer interface {
	AddFavorite(context.Context, *AddFavoriteRequest) (*AddFavoriteResponse, error)
	DelFavorite(context.Context, *DelFavoriteRequest) (*DelFavoriteResponse, error)
	GetVideoFavoriteCount(context.Context, *GetVideoFavoriteCountRequest) (*GetVideoFavoriteCountResponse, error)
	GetUserFavoriteCount(context.Context, *GetUserFavoriteCountRequest) (*GetUserFavoriteCountResponse, error)
	GetUserFavoritedCount(context.Context, *GetUserFavoritedCountRequest) (*GetUserFavoritedCountResponse, error)
	IsFavorite(context.Context, *IsFavoriteRequest) (*IsFavoriteResponse, error)
	GetFavoriteVideoIdList(context.Context, *GetFavoriteVideoIdListRequest) (*GetFavoriteVideoListIdResponse, error)
	AddCollection(context.Context, *AddCollectionRequest) (*AddCollectionResponse, error)
	DelCollection(context.Context, *DelCollectionRequest) (*DelCollectionResponse, error)
	GetUserCollectionList(context.Context, *GetUserCollectionIdListRequest) (*GetUserCollectionIdListResponse, error)
	GetUserCollectionCount(context.Context, *GetUserCollectionCountRequest) (*GetUserCollectionCountResponse, error)
	GetVideoCollectionCount(context.Context, *GetVideoCollectionCountRequest) (*GetVideoCollectionCountResponse, error)
	IsCollection(context.Context, *IsCollectionRequest) (*IsCollectionResponse, error)
	mustEmbedUnimplementedFavoriteRpcServer()
}

// UnimplementedFavoriteRpcServer must be embedded to have forward compatible implementations.
type UnimplementedFavoriteRpcServer struct {
}

func (UnimplementedFavoriteRpcServer) AddFavorite(context.Context, *AddFavoriteRequest) (*AddFavoriteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddFavorite not implemented")
}
func (UnimplementedFavoriteRpcServer) DelFavorite(context.Context, *DelFavoriteRequest) (*DelFavoriteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelFavorite not implemented")
}
func (UnimplementedFavoriteRpcServer) GetVideoFavoriteCount(context.Context, *GetVideoFavoriteCountRequest) (*GetVideoFavoriteCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVideoFavoriteCount not implemented")
}
func (UnimplementedFavoriteRpcServer) GetUserFavoriteCount(context.Context, *GetUserFavoriteCountRequest) (*GetUserFavoriteCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserFavoriteCount not implemented")
}
func (UnimplementedFavoriteRpcServer) GetUserFavoritedCount(context.Context, *GetUserFavoritedCountRequest) (*GetUserFavoritedCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserFavoritedCount not implemented")
}
func (UnimplementedFavoriteRpcServer) IsFavorite(context.Context, *IsFavoriteRequest) (*IsFavoriteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsFavorite not implemented")
}
func (UnimplementedFavoriteRpcServer) GetFavoriteVideoIdList(context.Context, *GetFavoriteVideoIdListRequest) (*GetFavoriteVideoListIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFavoriteVideoIdList not implemented")
}
func (UnimplementedFavoriteRpcServer) AddCollection(context.Context, *AddCollectionRequest) (*AddCollectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCollection not implemented")
}
func (UnimplementedFavoriteRpcServer) DelCollection(context.Context, *DelCollectionRequest) (*DelCollectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelCollection not implemented")
}
func (UnimplementedFavoriteRpcServer) GetUserCollectionList(context.Context, *GetUserCollectionIdListRequest) (*GetUserCollectionIdListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserCollectionList not implemented")
}
func (UnimplementedFavoriteRpcServer) GetUserCollectionCount(context.Context, *GetUserCollectionCountRequest) (*GetUserCollectionCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserCollectionCount not implemented")
}
func (UnimplementedFavoriteRpcServer) GetVideoCollectionCount(context.Context, *GetVideoCollectionCountRequest) (*GetVideoCollectionCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVideoCollectionCount not implemented")
}
func (UnimplementedFavoriteRpcServer) IsCollection(context.Context, *IsCollectionRequest) (*IsCollectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsCollection not implemented")
}
func (UnimplementedFavoriteRpcServer) mustEmbedUnimplementedFavoriteRpcServer() {}

// UnsafeFavoriteRpcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FavoriteRpcServer will
// result in compilation errors.
type UnsafeFavoriteRpcServer interface {
	mustEmbedUnimplementedFavoriteRpcServer()
}

func RegisterFavoriteRpcServer(s grpc.ServiceRegistrar, srv FavoriteRpcServer) {
	s.RegisterService(&FavoriteRpc_ServiceDesc, srv)
}

func _FavoriteRpc_AddFavorite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddFavoriteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FavoriteRpcServer).AddFavorite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FavoriteRpc_AddFavorite_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FavoriteRpcServer).AddFavorite(ctx, req.(*AddFavoriteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FavoriteRpc_DelFavorite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelFavoriteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FavoriteRpcServer).DelFavorite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FavoriteRpc_DelFavorite_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FavoriteRpcServer).DelFavorite(ctx, req.(*DelFavoriteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FavoriteRpc_GetVideoFavoriteCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVideoFavoriteCountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FavoriteRpcServer).GetVideoFavoriteCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FavoriteRpc_GetVideoFavoriteCount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FavoriteRpcServer).GetVideoFavoriteCount(ctx, req.(*GetVideoFavoriteCountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FavoriteRpc_GetUserFavoriteCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserFavoriteCountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FavoriteRpcServer).GetUserFavoriteCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FavoriteRpc_GetUserFavoriteCount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FavoriteRpcServer).GetUserFavoriteCount(ctx, req.(*GetUserFavoriteCountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FavoriteRpc_GetUserFavoritedCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserFavoritedCountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FavoriteRpcServer).GetUserFavoritedCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FavoriteRpc_GetUserFavoritedCount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FavoriteRpcServer).GetUserFavoritedCount(ctx, req.(*GetUserFavoritedCountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FavoriteRpc_IsFavorite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsFavoriteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FavoriteRpcServer).IsFavorite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FavoriteRpc_IsFavorite_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FavoriteRpcServer).IsFavorite(ctx, req.(*IsFavoriteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FavoriteRpc_GetFavoriteVideoIdList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFavoriteVideoIdListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FavoriteRpcServer).GetFavoriteVideoIdList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FavoriteRpc_GetFavoriteVideoIdList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FavoriteRpcServer).GetFavoriteVideoIdList(ctx, req.(*GetFavoriteVideoIdListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FavoriteRpc_AddCollection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCollectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FavoriteRpcServer).AddCollection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FavoriteRpc_AddCollection_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FavoriteRpcServer).AddCollection(ctx, req.(*AddCollectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FavoriteRpc_DelCollection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelCollectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FavoriteRpcServer).DelCollection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FavoriteRpc_DelCollection_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FavoriteRpcServer).DelCollection(ctx, req.(*DelCollectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FavoriteRpc_GetUserCollectionList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserCollectionIdListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FavoriteRpcServer).GetUserCollectionList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FavoriteRpc_GetUserCollectionList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FavoriteRpcServer).GetUserCollectionList(ctx, req.(*GetUserCollectionIdListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FavoriteRpc_GetUserCollectionCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserCollectionCountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FavoriteRpcServer).GetUserCollectionCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FavoriteRpc_GetUserCollectionCount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FavoriteRpcServer).GetUserCollectionCount(ctx, req.(*GetUserCollectionCountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FavoriteRpc_GetVideoCollectionCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVideoCollectionCountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FavoriteRpcServer).GetVideoCollectionCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FavoriteRpc_GetVideoCollectionCount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FavoriteRpcServer).GetVideoCollectionCount(ctx, req.(*GetVideoCollectionCountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FavoriteRpc_IsCollection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsCollectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FavoriteRpcServer).IsCollection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FavoriteRpc_IsCollection_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FavoriteRpcServer).IsCollection(ctx, req.(*IsCollectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FavoriteRpc_ServiceDesc is the grpc.ServiceDesc for FavoriteRpc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FavoriteRpc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "favorite.FavoriteRpc",
	HandlerType: (*FavoriteRpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddFavorite",
			Handler:    _FavoriteRpc_AddFavorite_Handler,
		},
		{
			MethodName: "DelFavorite",
			Handler:    _FavoriteRpc_DelFavorite_Handler,
		},
		{
			MethodName: "GetVideoFavoriteCount",
			Handler:    _FavoriteRpc_GetVideoFavoriteCount_Handler,
		},
		{
			MethodName: "GetUserFavoriteCount",
			Handler:    _FavoriteRpc_GetUserFavoriteCount_Handler,
		},
		{
			MethodName: "GetUserFavoritedCount",
			Handler:    _FavoriteRpc_GetUserFavoritedCount_Handler,
		},
		{
			MethodName: "IsFavorite",
			Handler:    _FavoriteRpc_IsFavorite_Handler,
		},
		{
			MethodName: "GetFavoriteVideoIdList",
			Handler:    _FavoriteRpc_GetFavoriteVideoIdList_Handler,
		},
		{
			MethodName: "AddCollection",
			Handler:    _FavoriteRpc_AddCollection_Handler,
		},
		{
			MethodName: "DelCollection",
			Handler:    _FavoriteRpc_DelCollection_Handler,
		},
		{
			MethodName: "GetUserCollectionList",
			Handler:    _FavoriteRpc_GetUserCollectionList_Handler,
		},
		{
			MethodName: "GetUserCollectionCount",
			Handler:    _FavoriteRpc_GetUserCollectionCount_Handler,
		},
		{
			MethodName: "GetVideoCollectionCount",
			Handler:    _FavoriteRpc_GetVideoCollectionCount_Handler,
		},
		{
			MethodName: "IsCollection",
			Handler:    _FavoriteRpc_IsCollection_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "favorite.proto",
}
