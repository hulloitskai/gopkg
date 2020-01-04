package configutil

import (
	"os"
)

// GoEnvKey is the key for the environmant variable that describes the current
// environment (i.e. "production", "development").
const GoEnvKey = "GOENV"

// Well-known values for 'GOENV'.
const (
	EnvProduction  = "production"
	EnvDevelopment = "development"
)

// GoEnv reteurns the current environment value for 'GOENV'.
func GoEnv() string { return os.Getenv(GoEnvKey) }

// LookupGoEnv is like GetEnv, except it additionally returns true if 'GOENV'
// was set, and false otherwise.
func LookupGoEnv() (string, bool) { return os.LookupEnv(GoEnvKey) }
