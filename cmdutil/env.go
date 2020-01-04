package cmdutil

import "go.stevenxie.me/gopkg/configutil"

// LoadEnv forwards a definition.
func LoadEnv(filenames ...string) error {
	return configutil.LoadEnv(filenames...)
}

// MustLoadEnv is like LoadEnv, but panics upon error.
func MustLoadEnv(filenames ...string) { must(LoadEnv(filenames...)) }
