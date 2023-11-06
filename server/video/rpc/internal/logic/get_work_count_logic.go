package logic

import (
	"context"
	"douniu/common/consts"
	"douniu/common/errorx"
	"douniu/server/video/rpc/internal/svc"
	"douniu/server/video/rpc/types/pb"
	"fmt"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetWorkCountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetWorkCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetWorkCountLogic {
	return &GetWorkCountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetWorkCountLogic) GetWorkCount(in *pb.WorkCountReq) (*pb.WorkCountResp, error) {
	var key string
	key = consts.VideoEveryUserTimeScore + fmt.Sprint(in.UserId)
	count, err := l.svcCtx.VideoModel.GetZsetCount(l.ctx, l.svcCtx.RedisClient, key)
	if err != nil {
		logc.Error(l.ctx, err)
		return nil, errors.Wrapf(errorx.NewDefaultError("查询用户视频总数err: "+err.Error()), "查询用户视频总数err:  RegisterReq：%v", in)
	}

	return &pb.WorkCountResp{
		WorkCount: count,
	}, nil
}
