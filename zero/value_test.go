package zero_test

import (
	"testing"
	"unsafe"

	"go.stevenxie.me/gopkg/zero"
)

func TestStruct(t *testing.T) {
	empty := zero.Empty()

	const expected = 0
	if actual := unsafe.Sizeof(empty); actual > expected {
		t.Errorf(
			"Expected empty struct to have a size of %d, but got a size of %d",
			expected, actual,
		)
	}
}
