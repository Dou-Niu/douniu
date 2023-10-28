package handler

import (
	"douniu/server/common/response"
	"net/http"

	"douniu/server/user/api/internal/logic"
	"douniu/server/user/api/internal/svc"
	"douniu/server/user/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ForgetPasswordHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ResetPassword
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewForgetPasswordLogic(r.Context(), svcCtx)
		err := l.ForgetPassword(&req)
		response.Response(r, w, 0, err) //②
	}
}
