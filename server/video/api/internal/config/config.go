package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	JWTAuth struct {
		AccessSecret       string
		AccessTokenExpire  int64
		RefreshTokenExpire int64
	}
	UserRpcConf  zrpc.RpcClientConf
	VideoRpcConf zrpc.RpcClientConf
}
