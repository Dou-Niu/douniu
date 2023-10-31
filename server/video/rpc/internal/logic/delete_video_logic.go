package logic

import (
	"context"
	"douniu/server/common/consts"
	"douniu/server/common/errorx"
	"fmt"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/mr"

	"douniu/server/video/rpc/internal/svc"
	"douniu/server/video/rpc/types/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteVideoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteVideoLogic {
	return &DeleteVideoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteVideoLogic) DeleteVideo(in *pb.DeleteVideoReq) (*pb.CommonResp, error) {

	// 并发删除
	err := mr.Finish(func() error {
		//  MySQL
		err := l.svcCtx.VideoModel.Delete(l.ctx, in.VideoId)
		return err
	}, func() error {
		// 删除全部视频时序排序
		_, err := l.svcCtx.RedisClient.ZRem(l.ctx, consts.VideoTimeScore, in.VideoId).Result()
		return err
	}, func() error {
		// 删除全部视频热度排序
		_, err := l.svcCtx.RedisClient.ZRem(l.ctx, consts.VideoHotScore, in.VideoId).Result()
		return err
	}, func() error {
		// 删除用户所有视频时序排序
		_, err := l.svcCtx.RedisClient.ZRem(l.ctx, consts.VideoEveryUserTimeScore+fmt.Sprint(in.MeUserId), in.VideoId).Result()
		return err
	}, func() error {
		// 删除用户所有视频热度排序
		_, err := l.svcCtx.RedisClient.ZRem(l.ctx, consts.VideoEveryUserHotScore+fmt.Sprint(in.MeUserId), in.VideoId).Result()
		return err
	}, func() error {
		// 删除分区所有视频时序排序
		_, err := l.svcCtx.RedisClient.ZRem(l.ctx, consts.VideoPartitionTimeScore+fmt.Sprint(in.Partition), in.VideoId).Result()
		return err
	}, func() error {
		// 删除分区所有视频热度排序
		_, err := l.svcCtx.RedisClient.ZRem(l.ctx, consts.VideoPartitionHotScore+fmt.Sprint(in.Partition), in.VideoId).Result()
		return err
	})
	if err != nil {
		return nil, errors.Wrapf(errorx.NewDefaultError("视频排序规则sort非法输入"), "视频排序规则sort非法输入 Req：%v", in)

	}
	if err != nil {
		return nil, errors.Wrapf(errorx.NewDefaultError("删除视频失败"+err.Error()), "更改用户密码失败ResetPassword：%v", in)
	}
	return &pb.CommonResp{}, nil
}
