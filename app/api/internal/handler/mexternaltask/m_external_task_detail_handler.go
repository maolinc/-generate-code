package mexternaltask

import (
	"net/http"

	"github.com/maolinc/gencode/app/api/internal/logic/mexternaltask"
	"github.com/maolinc/gencode/app/api/internal/svc"
	"github.com/maolinc/gencode/app/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func MExternalTaskDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IdReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := mexternaltask.NewMExternalTaskDetailLogic(r.Context(), svcCtx)
		resp, err := l.MExternalTaskDetail(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
