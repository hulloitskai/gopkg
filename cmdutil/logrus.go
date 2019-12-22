package cmdutil

import (
	"io"
	"os"
	"strings"

	"github.com/dmksnnk/sentryhook"
	raven "github.com/getsentry/raven-go"
	"github.com/sirupsen/logrus"

	"go.stevenxie.me/gopkg/configutil"
)

// Logrus-related environment keys.
const (
	EnvLogrusLevel  = "LOGRUS_LEVEL"
	EnvLogrusFormat = "LOGRUS_FORMAT"
	EnvLogrusCaller = "LOGRUS_CALLER"
)

// Valid environment values for the key EnvLogrusFormat.
const (
	LogrusFormatJSON = "json"
	LogrusFormatText = "text"
)

// NewLogger creates an application-level logrus.Entry.
func NewLogger(opts ...LogrusOption) *logrus.Entry {
	// Build config.
	cfg := LogrusConfig{
		Output: os.Stdout,
		TextFormatter: logrus.TextFormatter{
			EnvironmentOverrideColors: true,
		},
	}
	for _, opt := range opts {
		opt(&cfg)
	}

	// Create logger.
	log := logrus.New()
	log.SetOutput(cfg.Output)

	// Set 'report caller' option.
	{
		report := os.Getenv(EnvLogrusCaller)
		switch strings.ToLower(report) {
		case "true", "1":
			log.SetReportCaller(true)
		}
	}

	// Set level.
	if l, ok := os.LookupEnv(EnvLogrusLevel); ok {
		level, err := logrus.ParseLevel(l)
		if err != nil {
			Fatalf("Invalid '%s'; %v\n", EnvLogrusLevel, err)
		}
		log.SetLevel(level)
	} else if configutil.GetGoEnv() == configutil.GoEnvDevelopment {
		log.SetLevel(logrus.DebugLevel)
	}

	// Set formatter.
	var formatter logrus.Formatter
	if format, ok := os.LookupEnv(EnvLogrusFormat); ok {
		switch format {
		case LogrusFormatText:
			formatter = &cfg.TextFormatter
		case LogrusFormatJSON:
			formatter = &cfg.JSONFormatter
		default:
			Fatalf("Invalid '%s'; unknown formatter '%s'.", EnvLogrusFormat, format)
		}
	} else {
		switch configutil.GetGoEnv() {
		case configutil.GoEnvProduction:
			formatter = &cfg.JSONFormatter
		default:
			formatter = &cfg.TextFormatter
		}
	}
	log.SetFormatter(formatter)

	// Integrate error reporting with Sentry.
	if r := cfg.Raven; r != nil {
		hook := sentryhook.New(r)
		hook.SetAsync(logrus.ErrorLevel)
		hook.SetSync(logrus.PanicLevel, logrus.FatalLevel)
		log.AddHook(hook)
	}

	// Return entry from logger.
	return logrus.NewEntry(log)
}

// WithSentryHook adds a Sentry reporting hook to a logrus.Logger.
func WithSentryHook(rc *raven.Client) LogrusOption {
	return func(cfg *LogrusConfig) { cfg.Raven = rc }
}

type (
	// A LogrusConfig configures a logrus.Logger.
	LogrusConfig struct {
		// Output-related options.
		Output        io.Writer
		TextFormatter logrus.TextFormatter
		JSONFormatter logrus.JSONFormatter

		// Extensions-related options.
		Raven *raven.Client
	}

	// A LogrusOption modifies a LogrusConfig.
	LogrusOption func(*LogrusConfig)
)
