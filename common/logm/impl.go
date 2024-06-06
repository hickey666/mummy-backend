package logm

import (
	"context"
	"encoding/base64"
	"net/http"
	"sync"

	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/trace"
)

var traceMap sync.Map

type LogM struct {
	fields                    []logx.LogField
	responseBool              bool
	requestBool               bool
	headerBool                bool
	bodyBool                  bool
	request, response, header any
	body                      []byte
}

func Load(ctx context.Context) *LogM {
	traceID := trace.TraceIDFromContext(ctx)
	if traceID == "" {
		logc.Errorw(ctx, "traceID is empty")
		return &LogM{}
	}
	value, ok := traceMap.Load(traceID)
	lm := &LogM{}
	if ok {
		lm = value.(*LogM)
	} else {
		traceMap.Store(traceID, lm)
	}
	return lm
}

// SetUserID 用户ID
func (m *LogM) SetUserID(uid int64) *LogM {
	m.Append("user_id", uid)
	return m
}

// SetError 错误日志
func (m *LogM) SetError(err error) *LogM {
	if err != nil {
		m.Append("error", err)
	}
	return m
}

// SetToken token
func (m *LogM) SetToken(token string) *LogM {
	m.AppendSensitive("token", token)
	return m
}

// SetRequestID 请求ID
func (m *LogM) SetRequestID(requestID string) *LogM {
	m.Append("request_id", requestID)
	return m
}

// OpenResponse 开启响应日志
func (m *LogM) OpenResponse() *LogM {
	m.responseBool = true
	return m
}

// CloseResponse 关闭响应日志
func (m *LogM) CloseResponse() *LogM {
	m.responseBool = false
	return m
}

// SetResponse 响应日志
func (m *LogM) SetResponse(resp any) *LogM {
	m.response = resp
	return m
}

// OpenRequest 开启请求日志
func (m *LogM) OpenRequest() *LogM {
	m.requestBool = true
	return m
}

// CloseRequest 关闭请求日志
func (m *LogM) CloseRequest() *LogM {
	m.requestBool = false
	return m
}

// SetRequest  请求日志
func (m *LogM) SetRequest(req any) *LogM {
	m.request = req
	return m
}

// OpenHeader 开启请求头日志
func (m *LogM) OpenHeader() *LogM {
	m.headerBool = true
	return m
}

// CloseHeader 关闭请求头日志
func (m *LogM) CloseHeader() *LogM {
	m.headerBool = false
	return m
}

// SetHeader 请求头日志
func (m *LogM) SetHeader(header any) *LogM {
	m.header = header
	return m
}

// SetBody 请求体
func (m *LogM) SetBody(body []byte) *LogM {
	m.body = body
	return m
}

// OpenBody 开启请求体日志
func (m *LogM) OpenBody() *LogM {
	m.bodyBool = true
	return m
}

// CloseBody 关闭请求体日志
func (m *LogM) CloseBody() *LogM {
	m.bodyBool = false
	return m
}

// Close 关闭日志
func Close(ctx context.Context) {
	traceID := trace.TraceIDFromContext(ctx)
	if traceID == "" {
		return
	}
	traceMap.Delete(traceID)
}

// Append 添加日志
func (m *LogM) Append(name string, val any) *LogM {
	m.fields = append(m.fields, logx.Field(name, val))
	return m
}

// AppendSensitive 添加敏感信息
func (m *LogM) AppendSensitive(name string, val string) *LogM {
	ln := len(val)
	switch {
	case ln > 15:
		val = val[:7] + "***" + val[ln-7:]
	case ln > 7:
		val = val[:3] + "***" + val[ln-3:]
	case ln > 5:
		val = val[:2] + "***" + val[ln-2:]
	default:
		val = "***"
	}
	m.fields = append(m.fields, logx.Field(name, val))
	return m
}

type Request struct {
	UserAgent string `json:"user_agent"`
	IP        string `json:"ip"`
	FullUrl   string `json:"full_url"`
	Duration  string `json:"duration"`
}

func Info(ctx context.Context, req Request) {
	traceID := trace.TraceIDFromContext(ctx)
	if traceID == "" {
		return
	}
	load, ok := traceMap.Load(traceID)
	if ok {
		m := load.(*LogM)
		if m.responseBool {
			m.Append("response", m.response)
		}
		if m.requestBool {
			m.Append("request", m.request)
		}
		// 如果开启了请求头日志，则使用请求头日志，否则使用UserAgent
		if m.headerBool {
			m.Append("header", m.header)
		} else {
			m.Append("UserAgent", req.UserAgent)
		}
		if m.bodyBool {
			m.Append("body", base64.StdEncoding.EncodeToString(m.body))
		}
		data := m.fields
		data = append(data, logx.Field("IP", req.IP), logx.Field("FullUrl", req.FullUrl), logx.Field("Duration", req.Duration))
		logc.Infow(ctx, "AccessLog", data...)
		// 移除对象
		traceMap.Delete(traceID)
	}
}

func LoadResponse(r *http.Request, req, resp any, err error) {
	Load(r.Context()).SetError(err).SetRequest(req).SetResponse(resp)
}
