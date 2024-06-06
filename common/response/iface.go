package response

type ErrorIface interface {
	Error() string
	GetMessage() string
	GetCode() int
}
