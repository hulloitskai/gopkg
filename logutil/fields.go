package logutil

import (
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

// AddComponent appends the component name to the logrus.Entry.
func AddComponent(e *logrus.Entry, v zero.Interface) *logrus.Entry {
	name := name.OfType(v)
	if _, ok := e.Data[ComponentKey]; !ok {
		return e.WithField(ComponentKey, name)
	}

	// Component name exists, so append current component to it.
	var (
		field = e.Data[ComponentKey]
		cs    []string
	)
	if c, ok := field.(string); ok {
		cs = append(cs, c)
	} else if cs, ok = field.([]string); !ok {
		panic(errors.Newf(
			"logutil: entry contains component field of unknown type '%T'",
			field,
		))
	}
	cs = append(cs, name)
	return e.WithField(ComponentKey, cs)
}

// WithMethod adds the name of the method v to the logrus.Entry.
func WithMethod(e *logrus.Entry, v zero.Interface) *logrus.Entry {
	return e.WithField(MethodKey, name.OfMethod(v))
}
