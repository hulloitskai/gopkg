package mapstructx

import (
	"reflect"

	"github.com/cockroachdb/errors"
	"github.com/spf13/viper"

	"github.com/mitchellh/mapstructure"
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

// A DecoderConfigOption modifies a mapstructure.DecoderConfig.
type DecoderConfigOption = viper.DecoderConfigOption

// Unmarshal unmarshals the data from src into the value pointed to by dst.
func Unmarshal(
	src zero.Interface,
	dst zero.Interface,
	opts ...DecoderConfigOption,
) error {
	cfg := &mapstructure.DecoderConfig{
		Result:           dst,
		WeaklyTypedInput: true,
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			mapstructure.StringToTimeDurationHookFunc(),
			mapstructure.StringToSliceHookFunc(","),
			UnmarshalerHookFunc(),
		),
	}
	for _, opt := range opts {
		opt(cfg)
	}
	dec, err := mapstructure.NewDecoder(cfg)
	if err != nil {
		return errors.Wrap(err, "mapstructx: create decoder")
	}
	return dec.Decode(src)
}

// WithDecodeHooks creates a DecoderConfigOption that attaches appends to the
// existing hooks on the DecoderConfig.
func WithDecodeHooks(hooks ...mapstructure.DecodeHookFunc) DecoderConfigOption {
	return func(cfg *mapstructure.DecoderConfig) {
		var existing []mapstructure.DecodeHookFunc
		if cfg.DecodeHook != nil {
			existing = []mapstructure.DecodeHookFunc{cfg.DecodeHook}
		}
		hooks = append(existing, hooks...)
		ReplaceDecodeHooks(hooks)(cfg)
	}
}

// ReplaceDecodeHooks creates a DecoderConfigOption that replaces the existing
// decode hooks on the DecoderConfig.
func ReplaceDecodeHooks(hooks ...mapstructure.DecodeHookFunc) DecoderConfigOption {
	return func(cfg *mapstructure.DecoderConfig) {
		if n := len(hooks); n > 1 {
			cfg.DecodeHook = mapstructure.ComposeDecodeHookFunc(hooks...)
		} else if n == 1 {
			cfg.DecodeHook = hooks[0]
		}
	}
}

// WithTagName configures a mapstructure.Decoder to use struct tags named name.
func WithTagName(name string) DecoderConfigOption {
	return func(cfg *mapstructure.DecoderConfig) { cfg.TagName = name }
}
