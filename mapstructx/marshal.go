package mapstructx

import (
	"github.com/cockroachdb/errors"
	"github.com/spf13/viper"

	"github.com/mitchellh/mapstructure"
	"go.stevenxie.me/gopkg/zero"
)

// A DecoderConfigOption modifies a mapstructure.DecoderConfig.
type DecoderConfigOption = viper.DecoderConfigOption

// Unmarshal unmarshals the data from src into the value pointed to by dst.
func Unmarshal(
	src zero.Interface,
	dst zero.Interface,
	opts ...DecoderConfigOption,
) error {
	cfg := mapstructure.DecoderConfig{
		Result:           dst,
		WeaklyTypedInput: true,
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			mapstructure.StringToTimeDurationHookFunc(),
			mapstructure.StringToSliceHookFunc(","),
			UnmarshalerHookFunc(),
		),
	}
	for _, opt := range opts {
		opt(&cfg)
	}
	dec, err := mapstructure.NewDecoder(&cfg)
	if err != nil {
		return errors.Wrap(err, "mapstructx: create decoder")
	}
	return dec.Decode(src)
}

// Marshal marshals v into a map[string]interface{}.
func Marshal(
	v zero.Interface,
	opts ...DecoderConfigOption,
) (map[string]zero.Interface, error) {
	var m map[string]zero.Interface
	if err := Unmarshal(v, &m, opts...); err != nil {
		return nil, err
	}
	return m, nil
}

// WithDecodeHook creates a DecoderConfigOption that attaches appends to the
// existing hooks on the DecoderConfig.
func WithDecodeHook(hook ...mapstructure.DecodeHookFunc) DecoderConfigOption {
	return func(cfg *mapstructure.DecoderConfig) {
		var existing []mapstructure.DecodeHookFunc
		if cfg.DecodeHook != nil {
			existing = []mapstructure.DecodeHookFunc{cfg.DecodeHook}
		}
		hook = append(existing, hook...)
		ReplaceDecodeHook(hook...)(cfg)
	}
}

// ReplaceDecodeHook creates a DecoderConfigOption that replaces the existing
// decode hooks on the DecoderConfig.
func ReplaceDecodeHook(hook ...mapstructure.DecodeHookFunc) DecoderConfigOption {
	return func(cfg *mapstructure.DecoderConfig) {
		if n := len(hook); n > 1 {
			cfg.DecodeHook = mapstructure.ComposeDecodeHookFunc(hook...)
		} else if n == 1 {
			cfg.DecodeHook = hook[0]
		}
	}
}

// WithTagName configures a mapstructure.Decoder to use struct tags named name.
func WithTagName(name string) DecoderConfigOption {
	return func(cfg *mapstructure.DecoderConfig) { cfg.TagName = name }
}
