package configutil

import (
	stderrs "errors"
	"io"
	"os"

	"go.stevenxie.me/gopkg/cmdutil"

	"github.com/cockroachdb/errors"
	"go.stevenxie.me/api/pkg/zero"
	"gopkg.in/yaml.v2"
)

// TryLoadConfig loads configuration data from the first existing file in
// filenames, and stores it in cfg.
func TryLoadConfig(cfg zero.Interface, filenames ...string) error {
	for _, name := range filenames {
		ok, err := checkFile(name)
		if err != nil {
			return errors.WithMessage(err, "configutil")
		}
		if !ok {
			continue
		}
		return errors.Wrapf(
			LoadConfig(cfg, name),
			"configutil: loading '%s'", name,
		)
	}
	return ErrConfigNotFound
}

// ErrConfigNotFound is returned by TryLoadConfig when no config files could
// be found at any of the possible locations.
var ErrConfigNotFound = stderrs.New("configutil: no config files were found " +
	"at the specified locations")

// MustLoadConfig is like TryLoadConfig, but panics when an error occurs.
func MustLoadConfig(cfg zero.Interface, filenames ...string) {
	cmdutil.Must(TryLoadConfig(cfg, filenames...))
}

// LoadConfig loads configuration data from a file into cfg.
func LoadConfig(cfg zero.Interface, filename string) error {
	// Open file.
	file, err := os.Open(filename)
	if err != nil {
		return errors.Wrap(err, "configutil: opening file")
	}
	defer file.Close()

	// Read file.
	if err = ReadConfig(cfg, file); err != nil {
		return err
	}

	// Close file.
	if err = file.Close(); err != nil {
		return errors.Wrap(err, "configutil: closing file")
	}
	return nil
}

// ReadConfig reads in YAML configuration data from r, and stores it in the
// value pointed to by r.
func ReadConfig(cfg zero.Interface, r io.Reader) error {
	return yaml.NewDecoder(r).Decode(cfg)
}
