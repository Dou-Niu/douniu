package handler

import (
	"douniu/server/common/response"
	"net/http"

	"douniu/server/comment/api/internal/logic"
	"douniu/server/comment/api/internal/svc"
	"douniu/server/comment/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CommentActionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CommentActionRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewCommentActionLogic(r.Context(), svcCtx)
		resp, err := l.CommentAction(&req)
		response.Response(r, w, resp, err)
	}
}
