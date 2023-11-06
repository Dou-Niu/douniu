package logic

import (
	"context"
	"douniu/common/batcher"
	errorx2 "douniu/common/errorx"
	"douniu/server/video/model"
	"encoding/json"
	"github.com/pkg/errors"
	"strconv"
	"time"

	"douniu/server/video/rpc/internal/svc"
	"douniu/server/video/rpc/types/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublishVideoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	batcher *batcher.Batcher
}

func NewPublishVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishVideoLogic {
	f := &PublishVideoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
	// batcher配置
	options := batcher.Options{
		5,
		100,
		100,
		2 * time.Second,
	}
	// 实现batcher
	b := batcher.New(options)
	b.Sharding = func(key string) int {
		pid, _ := strconv.ParseInt(key, 10, 64)
		return int(pid) % options.Worker
	}
	b.Do = func(ctx context.Context, val map[string][]interface{}) {
		var msgs []*model.Video
		for _, vs := range val {
			for _, v := range vs {
				msgs = append(msgs, v.(*model.Video))
			}
		}
		kd, err := json.Marshal(msgs)
		if err != nil {
			logx.Errorf("Batcher.Do json.Marshal msgs: %v error: %v", msgs, err)
		}
		if err = f.svcCtx.KqPusherClient.Push(string(kd)); err != nil {
			logx.Errorf("KafkaPusher.Push kd: %s error: %v", string(kd), err)
		}
	}
	f.batcher = b
	f.batcher.Start()
	return f
}

func (l *PublishVideoLogic) PublishVideo(in *pb.PublishVideoReq) (*pb.CommonResp, error) {
	v := model.Video{
		Id:        l.svcCtx.Snowflake.Generate().Int64(),
		UserId:    in.MeUserID,
		Title:     in.Title,
		PlayUrl:   in.VideoUrl,
		CoverUrl:  in.CoverUrl,
		Partition: in.Partition,
	}

	// kafka异步处理file元数据
	err := l.batcher.Add(strconv.FormatInt(in.MeUserID, 10), &v)
	if err != nil {
		return nil, errors.Wrapf(errorx2.NewCodeError(10000, errorx2.ErrKafkaUserFileMeta+err.Error()), "kafka异步UserFileMeta失败 err:%v", err)
	}
	return &pb.CommonResp{}, nil
}
