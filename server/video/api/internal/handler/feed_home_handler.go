package handler

import (
	"douniu/common/response"
	"net/http"

	"douniu/server/video/api/internal/logic"
	"douniu/server/video/api/internal/svc"
	"douniu/server/video/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func FeedHomeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FeedHomeReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewFeedHomeLogic(r.Context(), svcCtx)
		resp, err := l.FeedHome(&req)
		response.Response(r, w, resp, err) //②

	}
}
