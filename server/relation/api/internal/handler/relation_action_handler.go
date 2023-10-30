package handler

import (
	"douniu/server/common/response"
	"net/http"

	"douniu/server/relation/api/internal/logic"
	"douniu/server/relation/api/internal/svc"
	"douniu/server/relation/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func RelationActionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RelationActionRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewRelationActionLogic(r.Context(), svcCtx)
		resp, err := l.RelationAction(&req)
		response.Response(r, w, resp, err)
	}
}
