package gHandler

import (
	"errors"

	"github.com/KevinZonda/GinTemplate/logger"
	"github.com/sirupsen/logrus"
)

type Resp[T any] struct {
	Data  T
	Error string
	E     error
}

func NewResp[T any]() *Resp[T] {
	return &Resp[T]{}
}

func (r *Resp[T]) Fail(err error) *Resp[T] {
	r.E = err
	r.Error = errToStr(err)
	logger.StdLogger().WithFields(logrus.Fields{
		"error": err,
	}).Error("Fail")
	return r
}

func (r *Resp[T]) FailByStr(err string) *Resp[T] {
	r.Fail(errors.New(err))
	return r
}

func (r *Resp[T]) WithErrMsg(msg string) *Resp[T] {
	r.Error = msg
	return r
}

func (r *Resp[T]) Ok(data T) {
	r.Data = data
}

func (r *Resp[T]) Auto(data T, err error) {
	if err != nil {
		r.Fail(err)
	} else {
		r.Ok(data)
	}
}

func (r *Resp[T]) FailIfErrNotNil(err error) bool {
	if err != nil {
		r.Fail(err)
		return true
	}
	return false
}

func (r *Resp[T]) FailIfErrNotNilWithMsg(err error, msg string) bool {
	if err != nil {
		r.Fail(err).WithErrMsg(msg)
		return true
	}
	return false
}

func errToStr(err error) string {
	if err == nil {
		return ""
	}
	return err.Error()
}
