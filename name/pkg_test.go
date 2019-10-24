package name_test

import (
	"fmt"

	"go.stevenxie.me/gopkg/name"
)

func ExamplePkgPathOf() {
	fmt.Println(name.PkgPathOf((*SomeStruct)(nil)))
	fmt.Println(name.PkgPathOf(ExamplePkgPathOf))

	// Output:
	// go.stevenxie.me/gopkg/name_test
	// go.stevenxie.me/gopkg/name_test
}
