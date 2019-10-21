package configutil

import (
	"github.com/cockroachdb/errors"
	"github.com/joho/godotenv"
)

// LoadEnv attempts to loads environment variables from 'dotenv' files, as
// specified by filenames. If no filenames are given, it will attempt to read
// '.env' and '.env.local' in the current directory.
//
// It will attempt to load environment variables from all the files in
// filenames, and will skip missing files.
func LoadEnv(filenames ...string) error {
	// Configure default filenames.
	if len(filenames) == 0 {
		filenames = []string{".env", ".env.local"}
	}

	for _, name := range filenames {
		ok, err := checkFile(name)
		if err != nil {
			return errors.WithMessage(err, "configutil")
		}
		if !ok {
			continue
		}

		// Load file using 'godotenv'.
		if err := godotenv.Load(name); err != nil {
			return errors.Wrapf(err, "configutil: loading '%s' with dotenv", name)
		}
	}

	return nil
}

// MustLoadEnv is like LoadEnv, but panics when an error is encountered.
func MustLoadEnv(filenames ...string) {
	must(LoadEnv(filenames...))
}
