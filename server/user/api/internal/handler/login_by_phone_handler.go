package handler

import (
	"douniu/common/response"
	"net/http"

	"douniu/server/user/api/internal/logic"
	"douniu/server/user/api/internal/svc"
	"douniu/server/user/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func LoginByPhoneHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterOrLoginByPhoneReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewLoginByPhoneLogic(r.Context(), svcCtx)
		resp, err := l.LoginByPhone(&req)
		response.Response(r, w, resp, err) //②
	}
}
