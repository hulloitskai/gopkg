package zero

import (
	"io/ioutil"

	"github.com/sirupsen/logrus"
)

// Logger returns a no-op logrus.Logger.
func Logger() *logrus.Logger { return noopLogger }

// Entry returns a no-op logrus.Entry.
func Entry() *logrus.Entry { return noopEntry }

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
