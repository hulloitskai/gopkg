package name

import (
	"reflect"
	"strings"

	"github.com/cockroachdb/errors"
	"github.com/iancoleman/strcase"
	"go.stevenxie.me/gopkg/zero"
)

// SQLTableFor derives an SQL table name from a value's type name.
//
// It returns a snake-cased string containing the value's full type name
// (excluding the package domain).
func SQLTableFor(v zero.Interface) string {
	t := getElem(v)
	if t.Kind() == reflect.Func {
		panic(errors.New("name: kind may not be 'reflect.Func'"))
	}
	path := t.PkgPath()
	if i := strings.LastIndexByte(path, '.'); i >= 0 {
		path = path[i+1:]
		if j := strings.IndexByte(path, '/'); j >= 0 {
			path = path[j+1:]
		}
	}
	path = strings.ReplaceAll(path, "/", "_")
	return path + "_" + strcase.ToSnake(t.Name())
}
