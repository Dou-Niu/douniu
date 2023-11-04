package logic

import (
	"context"
	"douniu/server/common/consts"
	"strconv"

	"douniu/server/relation/rpc/internal/svc"
	"douniu/server/relation/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserFollowCountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserFollowCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserFollowCountLogic {
	return &GetUserFollowCountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserFollowCountLogic) GetUserFollowCount(in *pb.GetUserFollowCountRequest) (resp *pb.GetUserFollowCountResponse, err error) {
	resp = new(pb.GetUserFollowCountResponse)
	userIdStr := strconv.FormatInt(in.UserId, 10)
	count, err := l.svcCtx.RedisClient.ZcardCtx(l.ctx, consts.UserFollowPrefix+userIdStr)
	if err != nil {
		l.Errorf("redis scard err: %v", err)
		return nil, err
	}
	resp.Count = int64(count)

	return

	return
}
