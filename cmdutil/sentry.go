package cmdutil

import (
	"os"

	raven "github.com/getsentry/raven-go"
	sentry "github.com/getsentry/sentry-go"
	"go.stevenxie.me/gopkg/configutil"
)

// Raven and Sentry-related environment variables.
const (
	EnvSentryDSN = "SENTRY_DSN"
)

// NewSentry creates a new sentry.Client.
func NewSentry(opts ...SentryOption) *sentry.Client {
	clientOpts := sentryClientOptions(opts...)
	client, err := sentry.NewClient(clientOpts)
	if err != nil {
		Fatalf("Failed to build Sentry client: %v\n", err)
	}
	return client
}

// InitSentry initializes the current sentry.Hub.
func InitSentry(opts ...SentryOption) {
	clientOpts := sentryClientOptions(opts...)
	sentry.Init(clientOpts)
}

func sentryClientOptions(opts ...SentryOption) sentry.ClientOptions {
	var cfg SentryConfig
	for _, opt := range opts {
		opt(&cfg)
	}

	// Construct client.
	dsn, ok := os.LookupEnv(EnvSentryDSN)
	if !ok {
		Fatalf(
			"cmdutil: missing environment variable '%s'\n",
			EnvSentryDSN,
		)
	}
	clientOpts := sentry.ClientOptions{
		Dsn:     dsn,
		Release: cfg.Release,
	}
	if env, ok := configutil.LookupGoEnv(); ok {
		clientOpts.Environment = env
	}

	return clientOpts
}

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

// WithSentryRelease sets the release tag for a raven.Client.
func WithSentryRelease(release string) SentryOption {
	return func(cfg *SentryConfig) { cfg.Release = release }
}

type (
	// SentryConfig configures a raven.Client.
	SentryConfig struct {
		Release string
	}

	// A SentryOption modifies a RavenConfig.
	SentryOption func(*SentryConfig)
)
