package errors

import (
	"errors"
	"runtime"
)

type Error struct {
	baseError    error
	serviceError error
	fileLocation string
	lineNumber   int
}

func NewError[T Error](base error, service error) *T {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "error.go"
		line = 15
	}
	return &T{
		baseError:    base,
		serviceError: service,
		fileLocation: file,
		lineNumber:   line,
	}
}

type IError interface {
	Base() error
	Service() error
}

func (e *Error) Error() string {
	return errors.Join(e.baseError, e.serviceError).Error()
}

func (e *Error) Base() error {
	return e.baseError
}
func (e *Error) Service() error {
	return e.serviceError
}
