package errors

import "errors"

var EnvFileReadFailure = errors.New("Error loading .env file")
var EnvFileVarNotFound = errors.New("Environment variable at key not found")
