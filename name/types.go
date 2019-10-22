package name

import (
	"errors"
	"reflect"
	"runtime"
	"strings"

	"go.stevenxie.me/gopkg/zero"
)

// OfType returns the type name of a value.
func OfType(v zero.Interface) string {
	t := getType(v)
	if t.Kind() == reflect.Func {
		return OfFunc(v)
	}
	return t.String()
}

// OfTypeFull returns the full type name of a value, including its import path.
func OfTypeFull(v zero.Interface) string {
	t := getType(v)
	if t.Kind() == reflect.Func {
		return OfFuncFull(v)
	}
	return t.PkgPath() + "." + t.Name()
}

// OfFunc returns the name of a function.
func OfFunc(v zero.Interface) string {
	name := OfFuncFull(v)
	tailIndex := strings.LastIndexByte(name, '/')
	if tailIndex == -1 {
		return name
	}
	return name[tailIndex+1:]
}

// OfFuncFull returns the full name of a function, within the context of its
// import path.
func OfFuncFull(v zero.Interface) string {
	f := runtime.FuncForPC(reflect.ValueOf(v).Pointer())
	return f.Name()
}

// OfMethod returns the name of a method on a struct or interface.
func OfMethod(v zero.Interface) string {
	name := OfFunc(v)
	tailIndex := strings.LastIndexByte(name, '.')
	if tailIndex == -1 {
		panic(errors.New("name: func name has no character '.'"))
	}
	return name[tailIndex+1:]
}

func getType(v zero.Interface) reflect.Type {
	t := reflect.TypeOf(v)
	switch t.Kind() {
	case reflect.Ptr, reflect.Interface:
		t = t.Elem()
	}
	return t
}
