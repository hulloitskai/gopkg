package cmdutil

import (
	"os"

	"go.stevenxie.me/gopkg/zero"
)

// FatalExitCode is the exit code that all Fatal* methods exit with after
// printing.
var FatalExitCode = 1

// Fatal is like Err, but exits with an error code after printing.
func Fatal(a ...zero.Interface) {
	if _, err := Err(a...); err != nil {
		panic(err)
	}
	os.Exit(FatalExitCode)
}

// Fatalln is like Errln, but exits with an error code after printing.
func Fatalln(a ...zero.Interface) {
	if _, err := Errln(a...); err != nil {
		panic(err)
	}
	os.Exit(FatalExitCode)
}

// Fatalf is like Errf, but exits with an error code after printing.
func Fatalf(format string, a ...zero.Interface) {
	if _, err := Errf(format, a...); err != nil {
		panic(err)
	}
	os.Exit(FatalExitCode)
}
