package logutil_test

import (
	"testing"

	"go.stevenxie.me/gopkg/logutil"
)

func TestWithComponent(t *testing.T) {
	var (
		e      = logutil.NoopEntry()
		expect = "package.Struct1"
	)

	e2 := logutil.WithComponent(e, expect)
	if c := e2.Data[logutil.ComponentKey]; c != expect {
		t.Errorf("e.Data[logutil.ComponentKey] = %v, want %s", c, expect)
	}

	if e2.Data[logutil.ComponentKey] == e.Data[logutil.ComponentKey] {
		t.Error("e2.Data[logutil.ComponentKey] == e.Data[logutil.ComponentKey], " +
			"expected different values")
	}

	component := "package.Struct2"
	expect = "package.Struct1::" + component
	e3 := logutil.WithComponent(e2, component)
	if c := e3.Data[logutil.ComponentKey]; c != expect {
		t.Errorf("e.Data[logutil.ComponentKey] = %v, want %s", c, expect)
	}
}
