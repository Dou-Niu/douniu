package logic

import (
	"context"
	"douniu/common/consts"
	"douniu/common/errorx"
	"douniu/common/utils"
	"douniu/server/video/rpc/types/pb"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"

	"douniu/server/video/api/internal/svc"
	"douniu/server/video/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublishVideoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPublishVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishVideoLogic {
	return &PublishVideoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PublishVideoLogic) PublishVideo(req *types.PublishVideoReq) error {
	err := utils.DefaultGetValidParams(l.ctx, req)
	if err != nil {
		return errors.Wrapf(errorx.NewCodeError(1, fmt.Sprintf("validate校验错误: %v", err)), "validate校验错误err :%v", err)
	}
	meId, err := l.ctx.Value(consts.UserId).(json.Number).Int64()
	_, err = l.svcCtx.VideoRpc.PublishVideo(l.ctx, &pb.PublishVideoReq{
		VideoId:   0,
		MeUserID:  meId,
		VideoUrl:  req.VideoUrl,
		CoverUrl:  req.CoverUrl,
		Partition: req.Partition,
		Title:     req.Title,
	})
	if err != nil {
		return errors.Wrapf(err, "req: %+v", req)
	}

	return nil
}
