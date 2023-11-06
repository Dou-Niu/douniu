package handler

import (
	"douniu/common/response"
	"net/http"

	"douniu/server/relation/api/internal/logic"
	"douniu/server/relation/api/internal/svc"
	"douniu/server/relation/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func RelationFollowListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RelationFollowListRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewRelationFollowListLogic(r.Context(), svcCtx)
		resp, err := l.RelationFollowList(&req)
		response.Response(r, w, resp, err)
	}
}
