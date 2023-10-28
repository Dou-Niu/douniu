package handler

import (
	"douniu/server/common/response"
	"net/http"

	"douniu/server/user/api/internal/logic"
	"douniu/server/user/api/internal/svc"
	"douniu/server/user/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func SendcodeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SendVerificationCodeReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewSendcodeLogic(r.Context(), svcCtx)
		resp, err := l.Sendcode(&req)
		response.Response(r, w, resp, err) //②
	}
}
