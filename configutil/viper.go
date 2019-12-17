package configutil

import (
	"fmt"

	"github.com/spf13/viper"
	"go.stevenxie.me/gopkg/mapstructx"
	"go.stevenxie.me/gopkg/zero"
)

// NewViper creates a new Viper that reads from a config file with the
// provided name, set in a namespace.
func NewViper(name, namespace string) *viper.Viper {
	v := viper.New()
	v.SetConfigName(name)
	v.AddConfigPath(".")
	v.AddConfigPath(fmt.Sprintf("/etc/%s", namespace))
	v.AddConfigPath(fmt.Sprintf("$HOME/.%s", namespace))
	return v
}

// UnmarshalViper unmarshals the configuration data in viper.Viper into the
// value pointed to by ptr.
func UnmarshalViper(
	v *viper.Viper,
	ptr zero.Interface,
	opts ...viper.DecoderConfigOption,
) error {
	opts = append(
		[]viper.DecoderConfigOption{
			mapstructx.WithDecodeHooks(mapstructx.UnmarshalerHookFunc()),
		},
		opts...,
	)
	return v.Unmarshal(ptr, opts...)
}
