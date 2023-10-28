package handler

import (
	"douniu/server/common/response"
	"net/http"

	"douniu/server/favorite/api/internal/logic"
	"douniu/server/favorite/api/internal/svc"
	"douniu/server/favorite/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func FavoriteListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FavoriteListRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewFavoriteListLogic(r.Context(), svcCtx)
		resp, err := l.FavoriteList(&req)
		response.Response(r, w, resp, err)
	}
}
