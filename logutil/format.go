package logutil

import (
	"fmt"
	"strconv"

	"github.com/cockroachdb/errors"
	"go.stevenxie.me/gopkg/mapstructx"
	"go.stevenxie.me/gopkg/zero"
)

// A Format represents an output formatting type for Logrus.
type Format uint8

var (
	_ fmt.Stringer           = (*Format)(nil)
	_ mapstructx.Unmarshaler = (*Format)(nil)
)

// The set of logrus formats.
const (
	JSONFormat Format = iota
	TextFormat
)

var formatNames = map[Format]string{
	JSONFormat: "JSON",
	TextFormat: "Text",
}

func (fmt Format) String() string {
	if name, ok := formatNames[fmt]; ok {
		return name
	}
	return "%!Format(" + strconv.FormatUint(uint64(fmt), 10) + ")"
}

// UnmarshalMap unmarshals fmt from a mapstructure.
func (fmt *Format) UnmarshalMap(v zero.Interface) error {
	name, ok := v.(string)
	if !ok {
		return errors.New("must be string")
	}
	for k, v := range formatNames {
		if v == name {
			*fmt = k
			return nil
		}
	}
	return errors.Newf("invalid Format '%s'", name)
}
