package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-courseware/courseware/api/internal/logic"
	"go-zero-courseware/courseware/api/internal/svc"
	"go-zero-courseware/courseware/api/internal/types"
)

func coursewareAddHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewCoursewareAddLogic(r.Context(), svcCtx)
		resp, err := l.CoursewareAdd(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
