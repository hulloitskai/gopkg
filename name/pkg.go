package name

import (
	"path"
	"reflect"
	"strings"

	"go.stevenxie.me/gopkg/zero"
)

// PkgPathOf returns the package path of v.
func PkgPathOf(v zero.Interface) string {
	t := getElem(v)
	if t.Kind() == reflect.Func {
		name := OfFuncFull(v)
		base := path.Base(name)
		base = base[:strings.IndexByte(base, '.')]
		return path.Dir(name) + "/" + base
	}
	return t.PkgPath()
}
