package configutil

import (
	"os"
)

// GoEnvKey is the key for the environmant variable that describes the current
// environment (i.e. "production", "development").
const GoEnvKey = "GOENV"

// Well-known values for 'GOENV'.
const (
	GoEnvProduction  = "production"
	GoEnvDevelopment = "development"
)

// GetGoEnv gets the current environment value for 'GOENV'.
func GetGoEnv() string { return os.Getenv(GoEnvKey) }

// LookupGoEnv is like GoGetEnv, except it also returns true if 'GOENV' was
// set, and false otherwise.
func LookupGoEnv() (string, bool) { return os.LookupEnv(GoEnvKey) }
