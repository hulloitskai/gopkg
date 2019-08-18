package configutil

import (
	"os"

	"github.com/cockroachdb/errors"
)

// checkFile ensures that a file exists and is not a directory.
func checkFile(name string) (ok bool, err error) {
	info, err := os.Stat(name)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, errors.Wrap(err, "checking file status")
	}
	if info.IsDir() {
		return false, errors.Newf("'%s' is a directory", name)
	}
	return true, nil
}
