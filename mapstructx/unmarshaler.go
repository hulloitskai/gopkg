package mapstructx

import (
	"reflect"

	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	"go.stevenxie.me/gopkg/zero"
)

// An Unmarshaler knows how to unmarshal itself from a map structure.
type Unmarshaler interface {
	UnmarshalMap(v zero.Interface) error
}

var unmarshalerType = reflect.TypeOf((*Unmarshaler)(nil)).Elem()

// UnmarshalerHookFunc returns a mapstructure.DecodeHookFunc that lets
// destination types that implement the Unmarshaler interface decode themselves.
func UnmarshalerHookFunc() mapstructure.DecodeHookFunc {
	return func(
		src reflect.Type,
		dst reflect.Type,
		val zero.Interface,
	) (zero.Interface, error) {
		if dst.Implements(unmarshalerType) {
			if dst.Kind() != reflect.Ptr {
				return nil, errors.New("mapstructx: target is not a pointer")
			}
			var (
				value        = reflect.New(dst.Elem())
				unmarshaller = value.Interface().(Unmarshaler)
			)
			if err := unmarshaller.UnmarshalMap(val); err != nil {
				return nil, errors.Wrap(err, "mapstructx: unmarshal map")
			}
			return unmarshaller, nil
		}
		return val, nil
	}
}
