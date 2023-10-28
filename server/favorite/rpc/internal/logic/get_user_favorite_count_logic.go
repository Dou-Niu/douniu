package logic

import (
	"context"
	"douniu/server/common/consts"
	"strconv"

	"douniu/server/favorite/rpc/internal/svc"
	"douniu/server/favorite/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserFavoriteCountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserFavoriteCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserFavoriteCountLogic {
	return &GetUserFavoriteCountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserFavoriteCountLogic) GetUserFavoriteCount(in *pb.GetUserFavoriteCountRequest) (resp *pb.GetUserFavoriteCountResponse, err error) {
	userIdStr := strconv.Itoa(int(in.UserId))
	resp = new(pb.GetUserFavoriteCountResponse)
	count, err := l.svcCtx.RedisClient.ZcardCtx(l.ctx, consts.UserFavoriteIdPrefix+userIdStr)
	if err != nil {
		l.Errorf("RedisClient ZcardCtx error: %v", err)
		return
	}
	resp.Count = int64(count)

	return
}
