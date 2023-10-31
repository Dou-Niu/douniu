package logic

import (
	"context"
	"douniu/server/common/consts"
	"douniu/server/video/rpc/types/pb"
	"encoding/json"
	"github.com/pkg/errors"

	"douniu/server/video/api/internal/svc"
	"douniu/server/video/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteVideoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteVideoLogic {
	return &DeleteVideoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteVideoLogic) DeleteVideo(req *types.DeleteVideoReq) error {
	meId, err := l.ctx.Value(consts.UserId).(json.Number).Int64()

	_, err = l.svcCtx.VideoRpc.DeleteVideo(l.ctx, &pb.DeleteVideoReq{
		VideoId:   req.VideoId,
		MeUserId:  meId,
		Partition: req.Partition,
	})
	if err != nil {
		return errors.Wrapf(err, "req: %+v", req)
	}

	return nil
}
