package mapstructx

import (
	"reflect"

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
			if src.Len() == 0 {
				return nil, nil
			}
			return reflect.ValueOf(val).Index(0).Interface(), nil
		}
		return val, nil
	}
}
