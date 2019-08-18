package zero

import (
	"io/ioutil"

	"github.com/sirupsen/logrus"
)

// Logger returns a no-op logrus.Logger.
func Logger() *logrus.Logger { return noopLogger }

var noopLogger *logrus.Logger

// Initialize noopLogger.
func init() {
	noopLogger = logrus.New()
	noopLogger.SetLevel(logrus.PanicLevel)
	noopLogger.SetOutput(ioutil.Discard)
}
