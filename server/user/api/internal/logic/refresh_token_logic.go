package logic

import (
	"context"
	"douniu/common/consts"
	"encoding/json"
	"github.com/golang-jwt/jwt/v4"
	"time"

	"douniu/server/user/api/internal/svc"
	"douniu/server/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RefreshTokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRefreshTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RefreshTokenLogic {
	return &RefreshTokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RefreshTokenLogic) RefreshToken() (resp *types.TokenResp, err error) {
	userId, err := l.ctx.Value(consts.UserId).(json.Number).Int64()
	// 生成accessToken
	auth := l.svcCtx.Config.JWTAuth
	now := time.Now().Unix()
	accessToken, _ := getJwtToken(auth.AccessSecret, now, auth.AccessTokenExpire, userId)
	// 生成refreshToken
	refreshToken, _ := getJwtToken(auth.AccessSecret, now, auth.RefreshTokenExpire, userId)
	return &types.TokenResp{
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
