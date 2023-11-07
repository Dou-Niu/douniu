package logic

import (
	"context"
	"douniu/server/chat/rpc/chatrpc"
	"douniu/server/relation/rpc/internal/svc"
	"douniu/server/relation/rpc/pb"
	"douniu/server/user/rpc/userrpc"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/threading"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFriendListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFriendListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFriendListLogic {
	return &GetFriendListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFriendListLogic) GetFriendList(in *pb.GetFriendListRequest) (resp *pb.GetFriendListResponse, err error) {
	userIdList, err := l.svcCtx.FollowModel.FindFriendIdList(l.ctx, in.ToUserId)
	if err != nil {
		l.Errorf("find friend id list err: %v", err)
		return nil, err
	}

	resp = new(pb.GetFriendListResponse)
	resp.UserList = make([]*pb.FriendUser, len(userIdList))
	ChatListResp, err := l.svcCtx.ChatRpc.MessageChatLast(l.ctx, &chatrpc.MessageChatLastRequest{
		FromUserId:   in.UserId,
		ToUserIdList: userIdList,
	})
	if err != nil {
		l.Errorf("chat rpc get message chat last err: %v", err)
		return nil, err
	}

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
			resp.UserList[i] = new(pb.FriendUser)
			_ = copier.Copy(resp.UserList[i], userInfoResp.Userinfo)
			resp.UserList[i].Message = ChatListResp.LastMessageList[i].Content
			resp.UserList[i].MsgType = ChatListResp.LastMessageList[i].MsgType
		})

	}

	group.Wait()

	if ResErr != nil {
		return nil, ResErr
	}

	return
}
