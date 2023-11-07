package handler

import (
	"douniu/common/response"
	"net/http"

	"douniu/server/user/api/internal/logic"
	"douniu/server/user/api/internal/svc"
	"douniu/server/user/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ModifyUserInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ModifyUserInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewModifyUserInfoLogic(r.Context(), svcCtx)
		err := l.ModifyUserInfo(&req)
		response.Response(r, w, 0, err) //②

	}
}
