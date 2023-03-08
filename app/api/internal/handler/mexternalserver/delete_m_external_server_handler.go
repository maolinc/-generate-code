package mexternalserver

import (
	"net/http"

	"github.com/maolinc/gencode/app/api/internal/logic/mexternalserver"
	"github.com/maolinc/gencode/app/api/internal/svc"
	"github.com/maolinc/gencode/app/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DeleteMExternalServerHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IdsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := mexternalserver.NewDeleteMExternalServerLogic(r.Context(), svcCtx)
		resp, err := l.DeleteMExternalServer(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
