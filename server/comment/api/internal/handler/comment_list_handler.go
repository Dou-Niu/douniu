package handler

import (
	"douniu/server/common/response"
	"net/http"

	"douniu/server/comment/api/internal/logic"
	"douniu/server/comment/api/internal/svc"
	"douniu/server/comment/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CommentListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CommentListRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewCommentListLogic(r.Context(), svcCtx)
		resp, err := l.CommentList(&req)
		response.Response(r, w, resp, err)
	}
}
