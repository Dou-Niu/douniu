package logic

import (
	"context"
	"douniu/server/comment/rpc/commentrpc"
	"douniu/server/user/rpc/userrpc"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/threading"

	"douniu/server/comment/rpc/internal/svc"
	"douniu/server/comment/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommentDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCommentDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentDetailLogic {
	return &GetCommentDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCommentDetailLogic) GetCommentDetail(in *pb.GetCommentDetailRequest) (resp *pb.GetCommentDetailResponse, err error) {
	resp = new(pb.GetCommentDetailResponse)
	subList, err := l.svcCtx.CommentModel.FindSub(l.ctx, in.CommentId)

	if err != nil {
		l.Errorf("CommentModel FindSub error：%v", err)
		return
	}
	resp.CommentList = make([]*pb.Comment, 0, len(subList))
	_ = copier.Copy(&resp.CommentList, subList)

	// 获取每一个评论的用户信息，这里使用协程组来并发获取
	group := threading.NewRoutineGroup()
	for i := 0; i < len(subList); i++ {
		ii := i
		group.RunSafe(func() {
			resp.CommentList[ii].User = new(commentrpc.User)
			userInfoResp, ierr := l.svcCtx.UserRpc.GetUserInfo(l.ctx, &userrpc.UserInfoReq{
				UserId:     subList[ii].UserId,
				FromUserId: in.UserId,
			})

			if err != nil {
				l.Errorf("Get user info error: %v", err)
				err = ierr
				return
			}

			_ = copier.Copy(resp.CommentList[ii].User, userInfoResp.Userinfo)
			//resp.CommentList[ii].User.Id = in.UserId
			//resp.CommentList[ii].User.Name = "test"
			//resp.CommentList[ii].User.Avatar = "https://avatars.githubusercontent.com/u/1967184?v=4"
		})
	}
	group.Wait()

	if err != nil {
		return nil, err
	}

	return
}
