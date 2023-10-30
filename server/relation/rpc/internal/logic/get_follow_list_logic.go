package logic

import (
	"context"
	"douniu/server/common/consts"
	"douniu/server/user/rpc/userrpc"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/threading"
	"strconv"
	"time"

	"douniu/server/relation/rpc/internal/svc"
	"douniu/server/relation/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFollowListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFollowListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFollowListLogic {
	return &GetFollowListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFollowListLogic) GetFollowList(in *pb.GetFollowListRequest) (resp *pb.GetFollowListResponse, err error) {
	toUserIdStr := strconv.FormatInt(in.ToUserId, 10)
	userIdListStr, err := l.svcCtx.RedisClient.ZrevrangebyscoreWithScoresAndLimitCtx(l.ctx, consts.UserFollowPrefix+toUserIdStr, 0, time.Now().Unix(), int(in.PageNum), consts.DefaultSize)
	if err != nil {
		l.Errorf("redis ZrevrangebyscoreWithScoresAndLimitCtx err: %v", err)
		return nil, err
	}
	userIdList := make([]int64, len(userIdListStr))
	for i, v := range userIdListStr {
		userIdList[i], _ = strconv.ParseInt(v.Key, 10, 64)
	}
	resp = new(pb.GetFollowListResponse)
	resp.UserList = make([]*pb.User, len(userIdList))

	group := threading.NewRoutineGroup()
	var ResErr error

	for i, userId := range userIdList {
		i, userId := i, userId
		group.RunSafe(func() {
			userInfoResp, err := l.svcCtx.UserRpc.GetUserInfo(l.ctx, &userrpc.UserInfoReq{
				UserId:     userId,
				FromUserId: in.UserId,
			})
			if err != nil {
				l.Errorf("user rpc get user info err: %v", err)
				ResErr = err
				return
			}
			resp.UserList[i] = new(pb.User)
			_ = copier.Copy(resp.UserList[i], userInfoResp.Userinfo)
		})
	}

	group.Wait()

	if ResErr != nil {
		return nil, ResErr
	}

	return
}
