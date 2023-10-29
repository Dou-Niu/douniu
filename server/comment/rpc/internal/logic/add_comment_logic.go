package logic

import (
	"context"
	"douniu/server/comment/model"
	"douniu/server/comment/rpc/commentrpc"
	"douniu/server/comment/rpc/internal/svc"
	"douniu/server/comment/rpc/pb"
	"douniu/server/common/consts"
	"douniu/server/user/rpc/userrpc"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"strconv"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCommentLogic {
	return &AddCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddCommentLogic) AddComment(in *pb.AddCommentRequest) (resp *pb.AddCommentResponse, err error) {
	// 对评论内容进行敏感词过滤
	in.Content = l.svcCtx.SensitiveWordFilter.Filter(in.Content)
	if in.Content == "" {
		return nil, errors.New("评论内容不能为空")
	}
	// 保存评论
	comment := &model.Comment{
		Id:         l.svcCtx.Snowflake.Generate().Int64(),
		Content:    in.Content,
		VideoId:    in.VideoId,
		UserId:     in.UserId,
		ParentId:   in.ParentId,
		CreateTime: time.Now().Unix(),
	}
	//// 丢到kafka里异步落库
	//commentJson, _ := jsonx.MarshalToString(&comment)
	//err = l.svcCtx.KafkaPusher.Push(commentJson)
	//if err != nil {
	//	l.Errorf("Push comment error: %v", err)
	//	return
	//}

	_, err = l.svcCtx.CommentModel.Insert(l.ctx, comment)
	if err != nil {
		l.Errorf("Insert comment error: %v", err)
		return
	}

	if in.ParentId != 0 {
		err = l.svcCtx.CommentModel.IncrSubCount(l.ctx, in.ParentId)
		if err != nil {
			l.Errorf("IncrSubCount error: %v", err)
			return
		}
	}

	// 视频评论数+1
	_, err = l.svcCtx.RedisClient.Incr(consts.VideoCommentCountPrefix + strconv.Itoa(int(comment.VideoId)))
	if err != nil {
		logx.Errorf("MessageAction RedisClient.Incr error: %s", err.Error())
		return
	}

	//获取用户信息
	userInfoResp, err := l.svcCtx.UserRpc.GetUserInfo(l.ctx, &userrpc.UserInfoReq{
		UserId:     in.UserId,
		FromUserId: in.UserId,
	})
	if err != nil {
		l.Errorf("Get user info error: %v", err)
		return
	}
	resp = new(pb.AddCommentResponse)
	resp.Comment = new(pb.Comment)
	_ = copier.Copy(resp.Comment, comment)
	resp.Comment.User = new(commentrpc.User)
	_ = copier.Copy(resp.Comment.User, userInfoResp)

	//resp.Comment.User.Id = in.UserId
	//resp.Comment.User.Username = "test"
	//resp.Comment.User.Avatar = "https://static001.geekbang.org/account/avatar/00/19/61/0b/1c0b7f0d.jpg"

	return
}
