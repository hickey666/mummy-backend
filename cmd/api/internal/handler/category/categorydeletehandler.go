package category

import (
	"net/http"

	"mummy-storage/cmd/api/internal/logic/category"
	"mummy-storage/cmd/api/internal/svc"
	"mummy-storage/cmd/api/internal/types"
	"mummy-storage/common/logm"
	"mummy-storage/common/response"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func CategoryDeleteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CategoryDeleteRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := category.NewCategoryDeleteLogic(r.Context(), svcCtx)
		resp, err := l.CategoryDelete(&req)
		logm.LoadResponse(r, req, resp, err)
		response.Done(w, resp, err)
	}
}
