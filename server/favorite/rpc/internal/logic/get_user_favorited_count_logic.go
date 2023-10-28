package logic

import (
	"context"
	"douniu/server/common/consts"
	"strconv"

	"douniu/server/favorite/rpc/internal/svc"
	"douniu/server/favorite/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserFavoritedCountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserFavoritedCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserFavoritedCountLogic {
	return &GetUserFavoritedCountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserFavoritedCountLogic) GetUserFavoritedCount(in *pb.GetUserFavoritedCountRequest) (resp *pb.GetUserFavoritedCountResponse, err error) {
	userIdStr := strconv.Itoa(int(in.UserId))
	resp = new(pb.GetUserFavoritedCountResponse)
	countStr, err := l.svcCtx.RedisClient.GetCtx(l.ctx, consts.UserFavoritedCountPrefix+userIdStr)
	if err != nil {
		l.Errorf("RedisClient GetCtx error: %v", err)
		return
	}
	if countStr == "" {
		resp.Count = 0
		return
	}

	count, err := strconv.Atoi(countStr)
	if err != nil {
		l.Errorf("strconv.Atoi error: %v", err)
		return
	}
	resp.Count = int64(count)

	return
}
