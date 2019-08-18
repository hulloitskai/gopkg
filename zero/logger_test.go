package zero_test

import (
	"os"

	"github.com/sirupsen/logrus"
	"go.stevenxie.me/gopkg/zero"
)

func ExampleLogger() {
	formatter := &logrus.TextFormatter{DisableTimestamp: true}

	// Regular logger will produce output.
	log := logrus.New()
	log.SetOutput(os.Stdout)
	log.SetFormatter(formatter)
	log.WithField("kind", "default").Info("Oh-ho, who's this?")

	// No-op logger will not produce output.
	noopLog := zero.Logger()
	noopLog.SetFormatter(formatter)
	noopLog.WithField("kind", "noop").Info("Hey, it's a-me.")

	// Output:
	// level=info msg="Oh-ho, who's this?" kind=default
}
