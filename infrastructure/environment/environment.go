package environment

import (
	"errors"
	"flag"
	"os"
	"strings"
)

// AppEnvironment enum
type AppEnvironment string

// IsLocal method of AppEnvironment
func (env AppEnvironment) IsLocal() bool {
	return env == LOCAL
}

// IsDev method of AppEnvironment
func (env AppEnvironment) IsDev() bool {
	return env == DEV
}

// const declaration of environment
const (
	LOCAL AppEnvironment = "local"
	DEV                  = "development"
)

// declare of error getting environment
var (
	ErrEnvironmentNotFound = errors.New("environment not found")
	strEnv                 = ""
)

// GetEnvFlag func to get command line flag "env"
func getEnvFlag(name, value, usage string) string {
	var flagEnv string
	flag.StringVar(&flagEnv, name, value, usage)
	flag.Parse()
	return flagEnv
}

// FromOsEnv func
func FromOsEnv() (AppEnvironment, error) {
	if strEnv == "" {
		strEnv = strings.Trim(strings.ToLower(os.Getenv("APP_ENV")), " ")
		if flag.Lookup("env") == nil {
			strEnv = getEnvFlag("env", "", "env Mode project local, development...")
		}
	}

	switch strEnv {
	case "development":
		return DEV, nil
	}

	return LOCAL, nil
}
