package handler

import (
	"douniu/common/response"
	"net/http"

	"douniu/server/user/api/internal/logic"
	"douniu/server/user/api/internal/svc"
)

func RefreshTokenHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewRefreshTokenLogic(r.Context(), svcCtx)
		resp, err := l.RefreshToken()
		response.Response(r, w, resp, err) //②
	}
}
