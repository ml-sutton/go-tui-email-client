package utils

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/ml-sutton/go-tui-email-client/internal/errors"
)

type EnvironmentVariables struct {
	environmentVariables map[string]string
}

func ReadEnv() (*EnvironmentVariables, *errors.Error) {
	var envVars []string = []string{}
	var environmentVariables = make(map[string]string)
	err := godotenv.Load()
	if err != nil {
		var base = err
		var service = errors.EnvFileReadFailure
		return nil, errors.NewError(base, service)
	}
	for _, variable := range envVars {
		value, envReadError := loadField(variable)
		if envReadError != nil {
			environmentVariables[variable] = "ggit"
		}
		environmentVariables[variable] = value
	}
	return &EnvironmentVariables{
		environmentVariables: environmentVariables,
	}, nil
}
func (e *EnvironmentVariables) Get(key string) string {
	if value, ok := e.environmentVariables[key]; !ok {
		return ""
	} else {
		return value
	}
}

func loadField(environmentVariableKey string) (string, *errors.Error) {
	var varFromEnv string = os.Getenv(environmentVariableKey)
	if "" == varFromEnv {
		var base = fmt.Errorf("Failed to find environment variable at %s", environmentVariableKey)
		var service = errors.EnvFileVarNotFound
		return "", errors.NewError(base, service)
	}
	return varFromEnv, nil
}
