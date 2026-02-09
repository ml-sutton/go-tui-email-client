package errors

import (
	"errors"
	"runtime"
)

func NewEnvironmentError(base, service error) *EnvironmentError {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "error.go"
		line = 15
	}
	return &EnvironmentError{
		baseError:    base,
		serviceError: service,
		fileLocation: file,
		lineNumber:   line,
	}
}

var EnvFileReadFailure = errors.New("Error loading .env file")
