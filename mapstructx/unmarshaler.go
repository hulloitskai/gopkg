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
		ptr := dst
		if dst.Kind() != reflect.Ptr {
			ptr = reflect.PtrTo(dst)
			if ptr == src {
				return val, nil
			}
		}
		if ptr.Implements(unmarshalerType) {
			var (
				value        = reflect.New(ptr.Elem())
				unmarshaller = value.Interface().(Unmarshaler)
			)
			if err := unmarshaller.UnmarshalMap(val); err != nil {
				return nil, errors.Wrap(err, "mapstructx: unmarshal map")
			}
			if ptr != dst {
				return value.Elem().Interface(), nil
			}
			return unmarshaller, nil
		}
		return val, nil
	}
}
