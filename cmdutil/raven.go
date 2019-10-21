package cmdutil

import (
	"os"

	raven "github.com/getsentry/raven-go"
	"go.stevenxie.me/gopkg/configutil"
)

// Raven and Sentry-related environment variables.
const (
	EnvSentryDSN = "SENTRY_DSN"
)

// NewRaven creates a new raven.Client.
func NewRaven(opts ...RavenOption) *raven.Client {
	var cfg RavenConfig
	for _, opt := range opts {
		opt(&cfg)
	}

	rc, err := raven.New(os.Getenv(EnvSentryDSN))
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

// WithRavenRelease sets the release tag for a raven.Client.
func WithRavenRelease(release string) RavenOption {
	return func(cfg *RavenConfig) { cfg.Release = release }
}

type (
	// RavenConfig configures a raven.Client.
	RavenConfig struct {
		Release string
	}

	// A RavenOption modifies a RavenConfig.
	RavenOption func(*RavenConfig)
)
