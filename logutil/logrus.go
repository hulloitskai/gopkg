package logutil

import (
	"io"
	"os"

	"github.com/dmksnnk/sentryhook"
	raven "github.com/getsentry/raven-go"
	"github.com/sirupsen/logrus"
	"go.stevenxie.me/gopkg/configutil"
)

// NewLogger creates an application-level logrus.Entry.
func NewLogger(opts ...Option) (*logrus.Entry, error) {
	// Build default config.
	cfg := Config{
		Level:         logrus.DebugLevel,
		Output:        os.Stderr,
		Format:        TextFormat,
		TextFormatter: logrus.TextFormatter{EnvironmentOverrideColors: true},
	}

	// Vary config based on GOENV.
	if configutil.GoEnv() == configutil.EnvProduction {
		cfg.Format = JSONFormat
		cfg.Level = logrus.InfoLevel
	}

	// Modify config from options.
	for _, opt := range opts {
		opt(&cfg)
	}

	// Create and configure logger.
	log := logrus.New()
	log.SetOutput(cfg.Output)
	log.SetReportCaller(cfg.ReportCaller)
	log.SetLevel(cfg.Level)

	// Set formatter.
	{
		var fmt logrus.Formatter
		switch cfg.Format {
		case TextFormat:
			fmt = &cfg.TextFormatter
		case JSONFormat:
			fmt = &cfg.JSONFormatter
		}
		log.SetFormatter(fmt)
	}

	// Integrate error reporting with Sentry.
	if client := cfg.Raven; client != nil {
		hook := sentryhook.New(client)
		hook.SetAsync(logrus.ErrorLevel)
		hook.SetSync(logrus.PanicLevel, logrus.FatalLevel)
		log.AddHook(hook)
	}

	// Return entry from logger.
	return logrus.NewEntry(log), nil
}

// LogrusWithLevel configures a logrus.Logger to log at lvl.
func LogrusWithLevel(lvl logrus.Level) Option {
	return func(cfg *Config) { cfg.Level = lvl }
}

// LogrusWithLevelString is like LogrusWithLevel, but parses the level from lvl.
//
// It panics if lvl is an invalid logrus.Level.
func LogrusWithLevelString(lvl string) Option {
	return func(cfg *Config) {
		var err error
		if cfg.Level, err = logrus.ParseLevel(lvl); err != nil {
			panic(err)
		}
	}
}

// LogrusWithCaller configures a logrus.Logger to report callers.
func LogrusWithCaller(enable bool) Option {
	return func(cfg *Config) { cfg.ReportCaller = enable }
}

// LogrusWithSentry adds a Sentry reporting hook to a logrus.Logger.
func LogrusWithSentry(client *raven.Client) Option {
	return func(cfg *Config) { cfg.Raven = client }
}

type (
	// A Config configures a logrus.Logger.
	Config struct {
		// Logging options.
		Level        logrus.Level
		ReportCaller bool

		// Output options.
		Output        io.Writer
		Format        Format
		TextFormatter logrus.TextFormatter
		JSONFormatter logrus.JSONFormatter

		// Extension options.
		Raven *raven.Client
	}

	// A Option modifies a LogrusConfig.
	Option func(*Config)
)