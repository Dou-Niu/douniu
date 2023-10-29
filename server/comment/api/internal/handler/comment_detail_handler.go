package handler

import (
	"net/http"

	"douniu/server/comment/api/internal/logic"
	"douniu/server/comment/api/internal/svc"
	"douniu/server/comment/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CommentDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CommentDetailRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewCommentDetailLogic(r.Context(), svcCtx)
		resp, err := l.CommentDetail(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
