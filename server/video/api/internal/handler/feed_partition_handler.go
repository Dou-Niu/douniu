package handler

import (
	"douniu/common/response"
	"net/http"

	"douniu/server/video/api/internal/logic"
	"douniu/server/video/api/internal/svc"
	"douniu/server/video/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func FeedPartitionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FeedPartitionReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewFeedPartitionLogic(r.Context(), svcCtx)
		resp, err := l.FeedPartition(&req)
		response.Response(r, w, resp, err) //②

	}
}
