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

// ParseFormat parses a Format from the string fmt.
func ParseFormat(fmt string) (Format, error) {
	switch fmt {
	case "json", "JSON":
		return JSONFormat, nil
	case "text", "Text":
		return TextFormat, nil
	default:
		return UnknownFormat, errors.Newf("logutil: unknown Format '%s'", fmt)
	}
}

var (
	_ fmt.Stringer           = (*Format)(nil)
	_ mapstructx.Unmarshaler = (*Format)(nil)
)

// The set of logrus formats.
const (
	UnknownFormat Format = iota
	JSONFormat
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
func (fmt *Format) UnmarshalMap(v zero.Interface) (err error) {
	name, ok := v.(string)
	if !ok {
		return errors.New("must be string")
	}
	*fmt, err = ParseFormat(name)
	return err
}
