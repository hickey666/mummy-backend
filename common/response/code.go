package response

import (
	"fmt"
)

type ErrorX struct {
	Code    int
	Message string
}

// Error
func (c *ErrorX) Error() string {
	return fmt.Sprintf("code(%d),message(%s)", c.Code, c.Message)
}

func (c *ErrorX) GetMessage() string {
	return c.Message
}

func (c *ErrorX) GetCode() int {
	return c.Code
}

func NewError(code int, message string) *ErrorX {
	return &ErrorX{code, message}
}

var (
	ErrInvalidParams = NewError(100001, "参数错误")
	ErrInternalError = NewError(100002, "内部错误")
)
