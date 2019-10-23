package name_test

import (
	"fmt"

	"go.stevenxie.me/gopkg/name"
)

func ExampleSQLTableFor() {
	fmt.Println(name.SQLTableFor((*SomeStruct)(nil)))

	// Output: gopkg_name_test_some_struct
}
