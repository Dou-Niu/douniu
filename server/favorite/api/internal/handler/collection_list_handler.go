package handler

import (
	"douniu/common/response"
	"net/http"

	"douniu/server/favorite/api/internal/logic"
	"douniu/server/favorite/api/internal/svc"
	"douniu/server/favorite/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CollectionListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CollectioneListRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewCollectionListLogic(r.Context(), svcCtx)
		resp, err := l.CollectionList(&req)
		response.Response(r, w, resp, err)
	}
}
