package cmdutil

// Must panics if err is not nil.
//
// It is used to wrap side-effect-inducing function calls that return an error.
func Must(err error) {
	if err != nil {
		panic(err)
	}
}
