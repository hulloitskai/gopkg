package logutil_test

import (
	"testing"

	"go.stevenxie.me/gopkg/logutil"
)

func TestAppendComponent(t *testing.T) {
	var (
		e      = logutil.NoopEntry()
		expect = "package.Struct1"
	)

	logutil.AppendComponent(e, expect)
	if c := e.Data[logutil.ComponentKey]; c != expect {
		t.Errorf("e.Data[logutil.ComponentKey] = %v, want %s", c, expect)
	}

	component := "package.Struct2"
	expect = "package.Struct1::" + component
	logutil.AppendComponent(e, component)
	if c := e.Data[logutil.ComponentKey]; c != expect {
		t.Errorf("e.Data[logutil.ComponentKey] = %v, want %s", c, expect)
	}
}
