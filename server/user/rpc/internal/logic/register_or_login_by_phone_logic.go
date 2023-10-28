package logic

import (
	"context"
	"douniu/server/common/consts"
	"douniu/server/common/errorx"
	"douniu/server/common/utils"
	"douniu/server/user/model"
	"douniu/server/user/rpc/internal/svc"
	"douniu/server/user/rpc/types/pb"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterOrLoginByPhoneLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterOrLoginByPhoneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterOrLoginByPhoneLogic {
	return &RegisterOrLoginByPhoneLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 使用验证码进行手机号注册或登录
func (l *RegisterOrLoginByPhoneLogic) RegisterOrLoginByPhone(in *pb.RegisterOrLoginByPhoneReq) (*pb.RegisterOrLoginResp, error) {
	// 校验验证码
	err := l.svcCtx.UserModel.IsCodeVerify(l.ctx, l.svcCtx.RedisClient, in.Phone, in.VerificationCode)
	if err != nil {
		return nil, errors.Wrapf(errorx.NewDefaultError(err.Error()), err.Error()+"RegisterOrLoginByPhoneReq:", in)
	}
	// 先检查是否已经注册

	u, err := l.svcCtx.UserModel.FindOneByPhone(l.ctx, in.Phone)
	if err != nil && !errors.Is(err, sqlc.ErrNotFound) {
		return nil, errors.Wrapf(errorx.NewDefaultError("redis查询手机号是否已经注册时候失败,err:"+err.Error()), "redis查询手机号是否已经注册时候失败 RegisterOrLoginByPhoneReq：%v", in)
	}
	var userId int64
	if errors.Is(err, sqlc.ErrNotFound) {
		// 注册
		logc.Info(l.ctx, fmt.Sprintf("手机号为 %v 的新用户开始注册", in.Phone))
		userId = l.svcCtx.Snowflake.Generate().Int64()
		u := model.User{
			Id:              userId,
			Nickname:        "斗牛用户" + fmt.Sprint(userId),
			Phone:           in.Phone,
			Password:        utils.Md5Password(utils.GeneratePassword(11), l.svcCtx.Config.Salt),
			Avatar:          consts.DefaultAvatar,
			BackgroundImage: consts.DefaultBackgroundImage,
			Signature:       consts.DefaultSignature,
			CreateTime:      time.Now(),
			UpdateTime:      time.Now(),
			DeleteTime:      time.Now(),
		}
		_, err := l.svcCtx.UserModel.Insert(l.ctx, &u)
		if err != nil {
			return nil, errors.Wrapf(errorx.NewDefaultError("注册失败，userinfo 写入mysql错误"), "注册失败，userinfo 写入mysql错误 RegisterReq：%v", in)
		}
		// 向redis添加phone

	} else {
		// 登录的用户
		userId = u.Id
	}
	// 生成accessToken
	auth := l.svcCtx.Config.JWTAuth
	now := time.Now().Unix()
	accessToken, _ := getJwtToken(auth.AccessSecret, now, auth.AccessTokenExpire, userId)
	// 生成refreshToken
	refreshToken, _ := getJwtToken(auth.AccessSecret, now, auth.RefreshTokenExpire, userId)
	return &pb.RegisterOrLoginResp{
		UserId:       userId,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
func getJwtToken(secretKey string, iat, seconds, userID int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userID
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
