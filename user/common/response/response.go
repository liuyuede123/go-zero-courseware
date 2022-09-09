package response

import (
	"go-zero-courseware/user/common/xerr"
	"net/http"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"google.golang.org/grpc/status"
)

type Response struct {
	Code    uint32      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

//http返回
func HttpResult(r *http.Request, w http.ResponseWriter, resp interface{}, err error) {

	if err == nil {
		//成功返回
		r := &Response{
			Code:    0,
			Message: "success",
			Data:    resp,
		}
		httpx.WriteJson(w, http.StatusOK, r)
	} else {
		//错误返回
		errcode := uint32(500)
		errmsg := "服务器错误"

		causeErr := errors.Cause(err)                // err类型
		if e, ok := causeErr.(*xerr.CodeError); ok { //自定义错误类型
			//自定义CodeError
			errcode = e.GetErrCode()
			errmsg = e.GetErrMsg()
		} else {
			if gstatus, ok := status.FromError(causeErr); ok { // grpc err错误
				grpcCode := uint32(gstatus.Code())
				errcode = grpcCode
				errmsg = gstatus.Message()
			}
		}

		logx.WithContext(r.Context()).Errorf("【API-ERR】 : %+v ", err)

		httpx.WriteJson(w, http.StatusBadRequest, &Response{
			Code:    errcode,
			Message: errmsg,
			Data:    nil,
		})
	}
}

func JwtUnauthorizedResult(w http.ResponseWriter, r *http.Request, err error) {
	httpx.WriteJson(w, http.StatusUnauthorized, &Response{401, "鉴权失败", nil})
}
