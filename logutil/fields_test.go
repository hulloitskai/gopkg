package logutil_test

import (
	"os"

	"github.com/sirupsen/logrus"
	"go.stevenxie.me/gopkg/logutil"
	"go.stevenxie.me/gopkg/zero"
)

type (
	Component1 zero.Struct
	Component2 zero.Interface
)

func ExampleAddComponent() {
	l := logrus.New()
	l.SetOutput(os.Stdout)
	l.SetFormatter(_logrusFormatter)
	log := logrus.NewEntry(l)

	log1 := logutil.AddComponent(log, (*Component1)(nil))
	log1.Info("Added first component.")

	log2 := logutil.AddComponent(log1, (*Component2)(nil))
	log2.Info("Added second component.")

	log1.Info("First logger components unchanged.")

	// Output:
	// level=info msg="Added first component." component=logutil_test.Component1
	// level=info msg="Added second component." component="[logutil_test.Component1 logutil_test.Component2]"
	// level=info msg="First logger components unchanged." component=logutil_test.Component1
}

type SomeStruct struct {
	log *logrus.Entry
}

func (ss SomeStruct) LogWithMethod() {
	logutil.
		WithMethod(ss.log, SomeStruct.LogWithMethod).
		Info("Hello from method!")
}

func ExampleWithMethod() {
	l := logrus.New()
	l.SetOutput(os.Stdout)
	l.SetFormatter(_logrusFormatter)

	ss := SomeStruct{log: logrus.NewEntry(l)}
	ss.LogWithMethod()
	// Output:
	// level=info msg="Hello from method!" method=LogWithMethod
}
