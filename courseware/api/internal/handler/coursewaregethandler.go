package handler

import (
	"go-zero-courseware/courseware/common/response"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-courseware/courseware/api/internal/logic"
	"go-zero-courseware/courseware/api/internal/svc"
	"go-zero-courseware/courseware/api/internal/types"
)

func coursewareGetHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewCoursewareGetLogic(r.Context(), svcCtx)
		resp, err := l.CoursewareGet(&req)
		response.HttpResult(r, w, resp, err)
	}
}
