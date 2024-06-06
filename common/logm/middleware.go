package logm

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/zeromicro/go-zero/core/timex"
	"github.com/zeromicro/go-zero/core/utils"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// BeforeFunc 可在此方法内重新设置日志的一些属性
var BeforeFunc func(lm *LogM, r *http.Request)

func Middleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		lm := Load(r.Context()).OpenRequest().OpenResponse()
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			lm.Append("ReadBodyError", err.Error())
		}
		lm.SetBody(body)
		lm.SetHeader(r.Header)
		r.Body = ioutil.NopCloser(bytes.NewReader(body))
		if BeforeFunc != nil {
			BeforeFunc(lm, r)
		}
		timer := utils.NewElapsedTimer()
		next(w, r)
		duration := timer.Duration()
		req := Request{
			UserAgent: r.UserAgent(),
			IP:        httpx.GetRemoteAddr(r),
			FullUrl:   r.RequestURI,
			Duration:  timex.ReprOfDuration(duration),
		}
		Info(r.Context(), req)
	}
}
