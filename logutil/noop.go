package logutil

import (
	"io/ioutil"

	"github.com/sirupsen/logrus"
)

// NoopLogger returns a no-op logrus.NoopLogger.
func NoopLogger() *logrus.Logger { return noopLogger }

// NoopEntry returns a no-op logrus.NoopEntry.
func NoopEntry() *logrus.Entry { return noopEntry }

var (
	noopLogger *logrus.Logger
	noopEntry  *logrus.Entry
)

// Initialize noopLogger, noopEntry.
func init() {
	noopLogger = logrus.New()
	noopLogger.SetLevel(logrus.PanicLevel)
	noopLogger.SetOutput(ioutil.Discard)

	noopEntry = logrus.NewEntry(noopLogger)
}
