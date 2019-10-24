package logutil

import (
	"strings"

	"go.stevenxie.me/gopkg/name"

	"github.com/cockroachdb/errors"
	"github.com/sirupsen/logrus"
	"go.stevenxie.me/gopkg/zero"
)

const (
	// ComponentKey is the field key used to store the component name.
	ComponentKey = "component"

	// MethodKey is the field key used to store the methodname.
	MethodKey = "method"
)

const _componentSep = "::"

// WithComponent appends the type name of v (the component containing the
// logger) to a logrus.Entry.
func WithComponent(e *logrus.Entry, v zero.Interface) *logrus.Entry {
	name := name.OfType(v)
	if _, ok := e.Data[ComponentKey]; !ok {
		return e.WithField(ComponentKey, name)
	}

	// Extract existing components, if any.
	var components []string
	if field, ok := e.Data[ComponentKey]; ok {
		cstr, ok := field.(string)
		if !ok {
			panic(errors.Newf(
				"logutil: entry contains component field of unknown type '%T'",
				field,
			))
		}
		components = append(components, cstr)
	}

	// Append latest component name to
	components = append(components, name)
	return e.WithField(ComponentKey, strings.Join(components, _componentSep))
}

// WithMethod adds the name of the method v to the logrus.Entry.
func WithMethod(e *logrus.Entry, v zero.Interface) *logrus.Entry {
	return e.WithField(MethodKey, name.OfMethod(v))
}
