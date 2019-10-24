package cmdutil

import (
	"os"

	raven "github.com/getsentry/raven-go"
	"go.stevenxie.me/gopkg/configutil"
)

// NewRaven creates a new raven.Client.
func NewRaven(opts ...SentryOption) *raven.Client {
	var cfg SentryConfig
	for _, opt := range opts {
		opt(&cfg)
	}

	dsn, ok := os.LookupEnv(EnvSentryDSN)
	if !ok {
		Fatalf(
			"cmdutil: missing environment variable '%s'\n",
			EnvSentryDSN,
		)
	}
	rc, err := raven.New(dsn)
	if err != nil {
		Fatalf("Failed to build Raven client: %v\n", err)
	}

	// Configure client.
	if env, ok := configutil.LookupGoEnv(); ok {
		rc.SetEnvironment(env)
	}
	if r := cfg.Release; r != "" {
		rc.SetRelease(r)
	}
	return rc
}
