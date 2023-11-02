package logic

import (
	"context"
	"douniu/server/common/consts"
	"strconv"

	"douniu/server/favorite/rpc/internal/svc"
	"douniu/server/favorite/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserCollectionCountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserCollectionCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserCollectionCountLogic {
	return &GetUserCollectionCountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserCollectionCountLogic) GetUserCollectionCount(in *pb.GetUserCollectionCountRequest) (resp *pb.GetUserCollectionCountResponse, err error) {
	userIdStr := strconv.Itoa(int(in.UserId))
	resp = new(pb.GetUserCollectionCountResponse)
	count, err := l.svcCtx.RedisClient.ZcardCtx(l.ctx, consts.UserCollectPrefix+userIdStr)
	if err != nil {
		l.Errorf("RedisClient ZcardCtx error: %v", err)
		return
	}
	resp.Count = int64(count)

	return
}
