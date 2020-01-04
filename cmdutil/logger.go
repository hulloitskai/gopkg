package cmdutil

import (
	"github.com/sirupsen/logrus"
	"go.stevenxie.me/gopkg/logutil"
)

// NewLogger forwards a definition.
func NewLogger(opts ...logutil.Option) (*logrus.Entry, error) {
	return logutil.NewLogger(opts...)
}

// MustNewLogger is like NewLogger, but panics upon error.
func MustNewLogger(opts ...logutil.Option) *logrus.Entry {
	log, err := NewLogger(opts...)
	must(err)
	return log
}
