package name_test

import (
	"fmt"

	"go.stevenxie.me/gopkg/name"
	"go.stevenxie.me/gopkg/zero"
)

type SomeIface zero.Interface

type SomeStruct struct {
	SomeField string
}

func (ss SomeStruct) SomeMethod() {}

func ExampleOfType() {
	fmt.Println(name.OfType((*SomeStruct)(nil)))
	fmt.Println(name.OfType((*SomeIface)(nil)))

	// Output:
	// name_test.SomeStruct
	// name_test.SomeIface
}

func ExampleOfTypeFull() {
	fmt.Println(name.OfTypeFull((*SomeStruct)(nil)))
	fmt.Println(name.OfTypeFull((*SomeIface)(nil)))

	// Output:
	// go.stevenxie.me/gopkg/name_test.SomeStruct
	// go.stevenxie.me/gopkg/name_test.SomeIface
}

func ExampleOfFunc() {
	fmt.Println(name.OfFunc(SomeStruct.SomeMethod))
	fmt.Println(name.OfFunc(ExampleOfFunc))

	// Output:
	// name_test.SomeStruct.SomeMethod
	// name_test.ExampleOfFunc
}

func ExampleOfFuncFull() {
	fmt.Println(name.OfFuncFull(SomeStruct.SomeMethod))
	fmt.Println(name.OfFuncFull(ExampleOfFunc))

	// Output:
	// go.stevenxie.me/gopkg/name_test.SomeStruct.SomeMethod
	// go.stevenxie.me/gopkg/name_test.ExampleOfFunc
}

func ExampleOfMethod() {
	fmt.Println(name.OfMethod(SomeStruct.SomeMethod))

	// Output:
	// SomeMethod
}
