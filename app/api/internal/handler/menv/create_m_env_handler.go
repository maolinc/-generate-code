package menv

import (
	"net/http"

	"github.com/maolinc/gencode/app/api/internal/logic/menv"
	"github.com/maolinc/gencode/app/api/internal/svc"
	"github.com/maolinc/gencode/app/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CreateMEnvHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateMEnvReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := menv.NewCreateMEnvLogic(r.Context(), svcCtx)
		resp, err := l.CreateMEnv(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
