package cmdutil

import (
	"fmt"
	"os"

	"go.stevenxie.me/gopkg/zero"
)

// Err is like fmt.Print, but writes to standard error.
func Err(a ...zero.Interface) (n int, err error) {
	return fmt.Fprint(os.Stderr, a...)
}

// Errln is like fmt.Println, but writes to standard error.
func Errln(a ...zero.Interface) (n int, err error) {
	return fmt.Fprintln(os.Stderr, a...)
}

// Errf is like fmt.Printf, but writes to standard error.
func Errf(format string, a ...zero.Interface) (n int, err error) {
	return fmt.Fprintf(os.Stderr, format, a...)
}
