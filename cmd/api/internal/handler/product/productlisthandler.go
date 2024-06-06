package product

import (
	"net/http"

	"mummy-storage/cmd/api/internal/logic/product"
	"mummy-storage/cmd/api/internal/svc"
	"mummy-storage/cmd/api/internal/types"
	"mummy-storage/common/logm"
	"mummy-storage/common/response"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func ProductListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ProductListRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := product.NewProductListLogic(r.Context(), svcCtx)
		resp, err := l.ProductList(&req)
		logm.LoadResponse(r, req, resp, err)
		response.Done(w, resp, err)
	}
}
