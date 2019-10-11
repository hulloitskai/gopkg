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

	logutil.WithComponent(e, expect)
	if c := e.Data[logutil.ComponentKey]; c != expect {
		t.Errorf("e.Data[logutil.ComponentKey] = %v, want %s", c, expect)
	}

	component := "package.Struct2"
	expect = "package.Struct1::" + component
	logutil.WithComponent(e, component)
	if c := e.Data[logutil.ComponentKey]; c != expect {
		t.Errorf("e.Data[logutil.ComponentKey] = %v, want %s", c, expect)
	}
}
