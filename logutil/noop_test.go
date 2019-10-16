package logutil_test

import (
	"os"

	"github.com/sirupsen/logrus"
	"go.stevenxie.me/gopkg/logutil"
)

var _logrusFormatter = &logrus.TextFormatter{DisableTimestamp: true}

func ExampleNoopEntry() {
	// Regular logger will produce output.
	log := logrus.New()
	log.SetOutput(os.Stdout)
	log.SetFormatter(_logrusFormatter)
	log.WithField("kind", "default").Info("Oh-ho, who's this?")

	// No-op logger will not produce output.
	noopLog := logutil.NoopLogger()
	noopLog.SetFormatter(_logrusFormatter)
	noopLog.WithField("kind", "noop").Info("Hey, it's a-me.")

	// Output:
	// level=info msg="Oh-ho, who's this?" kind=default
}
