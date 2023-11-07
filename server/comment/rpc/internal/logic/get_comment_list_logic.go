package logic

import (
	"context"
	"douniu/server/comment/rpc/commentrpc"
	"douniu/server/comment/rpc/internal/svc"
	"douniu/server/comment/rpc/pb"
	"douniu/server/user/rpc/userrpc"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/threading"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommentListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentListLogic {
	return &GetCommentListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCommentListLogic) GetCommentList(in *pb.GetCommentListRequest) (resp *pb.GetCommentListResponse, err error) {
	commentList, err := l.svcCtx.CommentModel.FindByVideoId(l.ctx, in.VideoId, in.LastCommentId)
	if err != nil {
		l.Errorf("Get comment list error: %v", err)
		return
	}

	resp = new(pb.GetCommentListResponse)
	resp.CommentList = make([]*pb.Comment, len(commentList))
	_ = copier.Copy(&resp.CommentList, &commentList)

	// 获取每一个评论的用户信息，这里使用协程组来并发获取
	group := threading.NewRoutineGroup()
	for i := 0; i < len(commentList); i++ {
		ii := i
		group.RunSafe(func() {
			resp.CommentList[ii].User = new(commentrpc.User)
			userInfoResp, ierr := l.svcCtx.UserRpc.GetUserInfo(l.ctx, &userrpc.UserInfoReq{
				UserId:     commentList[ii].UserId,
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
