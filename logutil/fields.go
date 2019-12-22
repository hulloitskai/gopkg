package logutil

import (
	"path"

	"github.com/cockroachdb/errors"
	"go.stevenxie.me/gopkg/name"

	"github.com/sirupsen/logrus"
	"go.stevenxie.me/gopkg/zero"
)

const (
	// TypeKey is the field key used to store the logging object's type name.
	TypeKey = "type"

	// MethodKey is the field key used to store the method name where the logger
	// is being called.
	MethodKey = "method"

	// ComponentKey is the field key used to store the logging object's component
	// path.
	ComponentKey = "component"
)

// WithType adds type name of v to a logrus.Entry.
func WithType(e *logrus.Entry, v zero.Interface) *logrus.Entry {
	name := name.OfType(v)
	return e.WithField(TypeKey, name)
}

// WithMethod adds the name of the method v to the logrus.Entry.
func WithMethod(e *logrus.Entry, v zero.Interface) *logrus.Entry {
	return e.WithField(MethodKey, name.OfMethod(v))
}

// WithComponent appends the component to the component path field of the
// logrus.Entry.
func WithComponent(e *logrus.Entry, component string) *logrus.Entry {
	if _, ok := e.Data[ComponentKey]; !ok {
		return e.WithField(ComponentKey, component)
	}

	// Extract existing component elems, if any.
	var elems []string
	if component, ok := e.Data[ComponentKey]; ok {
		s, ok := component.(string)
		if !ok {
			panic(errors.Newf(
				"logutil: entry contains component field of unknown type '%T'",
				component,
			))
		}
		elems = append(elems, s)
	}

	// Append latest component name to
	elems = append(elems, component)
	return e.WithField(ComponentKey, path.Join(elems...))
}
