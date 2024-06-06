package errcode

import "mummy-storage/common/response"

var (
	ErrCheckeCycle = response.NewError(20001, "cycle格式错误")
	ErrTimeLayout  = response.NewError(20002, "时间格式错误")

	ErrRedis   = response.NewError(50001, "redis错误")
	ErrService = response.NewError(50002, "服务错误")
)
