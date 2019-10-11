package logutil

import (
	"strings"

	"github.com/sirupsen/logrus"
)

// ComponentKey is the field key used to store the component field.
const ComponentKey = "component"

// AppendComponent appends to the component field in the logrus.Entry.
//
// This is used to namespace log entries.
func AppendComponent(e *logrus.Entry, component ...string) {
	if c, ok := e.Data[ComponentKey].(string); ok {
		component = append([]string{c}, component...)
	}
	e.Data[ComponentKey] = strings.Join(component, "::")
}
