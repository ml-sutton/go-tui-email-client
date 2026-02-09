package utils

import (
	"github.com/joho/godotenv"
	"github.com/ml-sutton/go-tui-email-client/internal/errors"
)

type EnvironmentVariables struct {
}

func ReadENV() (*EnvironmentVariables, *errors.EnvironmentError) {
	err := godotenv.Load()
	if err != nil {
		var base = err
		var service = errors.EnvFileReadFailure
		return nil, errors.NewEnvironmentError(base, service)
	}

	return &EnvironmentVariables{}, nil
}
