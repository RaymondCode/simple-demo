package video

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"tikstart/http/internal/logic/video"
	"tikstart/http/internal/svc"
	"tikstart/http/internal/types"
)

func GetVideoListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetVideoListRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := video.NewGetVideoListLogic(r.Context(), svcCtx)
		resp, err := l.GetVideoList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
