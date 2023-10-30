package handler

import (
	"douniu/server/common/response"
	"net/http"

	"douniu/server/relation/api/internal/logic"
	"douniu/server/relation/api/internal/svc"
	"douniu/server/relation/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func RelationFriendListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FriendListRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewRelationFriendListLogic(r.Context(), svcCtx)
		resp, err := l.RelationFriendList(&req)
		response.Response(r, w, resp, err)
	}
}
