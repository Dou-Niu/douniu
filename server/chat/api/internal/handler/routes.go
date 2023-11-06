// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"douniu/server/chat/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/douniu/message/action",
				Handler: MessageActionHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/douniu/message/chat",
				Handler: MessageChatHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)
}
