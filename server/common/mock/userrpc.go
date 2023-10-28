package mock

import (
	"context"
	"douniu/server/user/rpc/userrpc"
	"google.golang.org/grpc"
)

type UserRpc struct {
}

func NewUserRpc() UserRpc {
	return UserRpc{}
}

func (u UserRpc) SendVerificationCode(ctx context.Context, in *userrpc.SendVerificationCodeReq, opts ...grpc.CallOption) (*userrpc.SendVerificationCodeResp, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserRpc) RegisterOrLoginByPhone(ctx context.Context, in *userrpc.RegisterOrLoginByPhoneReq, opts ...grpc.CallOption) (*userrpc.RegisterOrLoginResp, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserRpc) RegisterOrLoginByPassword(ctx context.Context, in *userrpc.RegisterOrLoginByPasswordReq, opts ...grpc.CallOption) (*userrpc.RegisterOrLoginResp, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserRpc) ForgetPassword(ctx context.Context, in *userrpc.ResetPassword, opts ...grpc.CallOption) (*userrpc.CommonResp, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserRpc) ChangePassword(ctx context.Context, in *userrpc.ResetPassword, opts ...grpc.CallOption) (*userrpc.CommonResp, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserRpc) GetUserInfo(ctx context.Context, in *userrpc.UserInfoReq, opts ...grpc.CallOption) (*userrpc.UserInfoResp, error) {
	//TODO implement me
	panic("implement me")
}
