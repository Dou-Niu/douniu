package handler

import (
	"douniu/server/common/response"
	"net/http"

	"douniu/server/relation/api/internal/logic"
	"douniu/server/relation/api/internal/svc"
	"douniu/server/relation/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func RelationFollowerListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RelationFollowerListRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewRelationFollowerListLogic(r.Context(), svcCtx)
		resp, err := l.RelationFollowerList(&req)
		response.Response(r, w, resp, err)
	}
}
