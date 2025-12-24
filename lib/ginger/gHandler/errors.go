package gHandler

const REQ_BIND_FAILED = "data binding failed"

type IHttpError interface {
	error
	Code() int
}

type HttpError struct {
	S string
	C int
}

func (e HttpError) Error() string {
	return e.S
}

func (e HttpError) Code() int {
	return e.C
}

func NewHttpError(s string, c int) IHttpError {
	return HttpError{S: s, C: c}
}
