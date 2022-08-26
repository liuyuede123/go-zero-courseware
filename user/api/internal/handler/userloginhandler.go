package handler

import (
	"go-zero-courseware/user/api/response"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-courseware/user/api/internal/logic"
	"go-zero-courseware/user/api/internal/svc"
	"go-zero-courseware/user/api/internal/types"
)

func userLoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewUserLoginLogic(r.Context(), svcCtx)
		resp, err := l.UserLogin(&req)
		response.Response(w, resp, err)
	}
}
