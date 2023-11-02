package logic

import (
	"context"
	"douniu/server/common/consts"
	"strconv"
	"time"

	"douniu/server/favorite/rpc/internal/svc"
	"douniu/server/favorite/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserCollectionListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserCollectionListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserCollectionListLogic {
	return &GetUserCollectionListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserCollectionListLogic) GetUserCollectionList(in *pb.GetUserCollectionListRequest) (resp *pb.GetUserCollectionListResponse, err error) {
	resp = new(pb.GetUserCollectionListResponse)
	//idListStr, err := l.svcCtx.RedisClient.ZrevrangeCtx(l.ctx, consts.UserFavoriteIdPrefix+strconv.Itoa(int(in.UserId)), 0, -1)
	idListStr, err := l.svcCtx.RedisClient.ZrevrangebyscoreWithScoresAndLimitCtx(l.ctx, consts.UserCollectPrefix+strconv.Itoa(int(in.UserId)), 0, time.Now().Unix(), int(in.PageNum), consts.DefaultSize)
	if err != nil {
		l.Errorf("RedisClient ZrangeCtx error: %v", err)
		return
	}

	idList := make([]int64, 0, len(idListStr))
	for _, idStr := range idListStr {
		id, err := strconv.Atoi(idStr.Key)
		if err != nil {
			l.Errorf("strconv.Atoi error: %v", err)
			return nil, err
		}
		idList = append(idList, int64(id))
	}
	resp.VideoIdList = idList

	return
}
