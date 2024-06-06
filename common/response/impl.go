package response

import (
	"net/http"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/rest/httpx"
)

type Body struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func Done(w http.ResponseWriter, resp interface{}, err error) {
	var body = Body{
		Data: resp,
		Msg:  "OK",
		Code: 0,
	}
	if err != nil {
		body.Code = -1
		causeErr := errors.Cause(err)
		body.Msg = causeErr.Error()
		if v, ok := causeErr.(ErrorIface); ok {
			body.Code = v.GetCode()
			body.Msg = v.GetMessage()
		}
	}
	httpx.OkJson(w, body)
}
