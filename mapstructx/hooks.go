package mapstructx

import (
	"reflect"

	"github.com/cockroachdb/errors"
	"github.com/mitchellh/mapstructure"
	"go.stevenxie.me/gopkg/zero"
)

// HCLHookFunc returns a mapstructure.DecodeHookFunc that helps unmarshal
// HCL configs by allowing weakly typed fields where slices are typically
// expected.
func HCLHookFunc() mapstructure.DecodeHookFunc {
	return func(
		src reflect.Type,
		dst reflect.Type,
		val zero.Interface,
	) (zero.Interface, error) {
		// If source is a slice and destination is not, continue with first element
		// of source slice.
		if (src.Kind() == reflect.Slice) && (dst.Kind() != reflect.Slice) {
			var (
				v = reflect.ValueOf(val)
				n = v.Len()
			)
			if n == 0 {
				return nil, nil
			}
			var err error
			if n > 1 {
				err = errors.New("mapstructx: slice contains more than one elem")
			}
			return v.Index(0).Interface(), err
		}
		return val, nil
	}
}
