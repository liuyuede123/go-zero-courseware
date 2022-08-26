package response

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

type Body struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func Response(w http.ResponseWriter, resp interface{}, err error) {
	var body Body
	if err != nil {
		body.Code = 5000
		body.Message = err.Error()
		body.Data = nil
	} else {
		body.Code = 0
		body.Message = "OK"
		body.Data = resp
	}
	httpx.OkJson(w, body)
}
