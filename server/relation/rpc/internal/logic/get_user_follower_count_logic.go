package logic

import (
	"context"
	"douniu/server/common/consts"
	"strconv"

	"douniu/server/relation/rpc/internal/svc"
	"douniu/server/relation/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserFollowerCountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserFollowerCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserFollowerCountLogic {
	return &GetUserFollowerCountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserFollowerCountLogic) GetUserFollowerCount(in *pb.GetUserFollowerCountRequest) (resp *pb.GetUserFollowerCountResponse, err error) {
	resp = new(pb.GetUserFollowerCountResponse)
	userIdStr := strconv.FormatInt(in.UserId, 10)
	count, err := l.svcCtx.RedisClient.ScardCtx(l.ctx, consts.UserFollowerPrefix+userIdStr)
	if err != nil {
		l.Errorf("redis scard err: %v", err)
		return nil, err
	}
	resp.Count = count

	return
}
