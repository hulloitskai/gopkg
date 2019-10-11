package logutil

import (
	"strings"

	"github.com/sirupsen/logrus"
)

// ComponentKey is the field key used to store the component field.
const ComponentKey = "component"

// WithComponent appends to the component field in the logrus.Entry.
//
// This is used to namespace log entries.
func WithComponent(e *logrus.Entry, component ...string) *logrus.Entry {
	if c, ok := e.Data[ComponentKey].(string); ok {
		component = append([]string{c}, component...)
	}
	return e.WithField(ComponentKey, strings.Join(component, "::"))
}
