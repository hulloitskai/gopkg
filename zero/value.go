package zero

// Empty returns an empty Struct.
func Empty() Struct { return Struct{} }

type (
	// Struct is a struct with no fields.
	//
	// It holds no information, and has a size of zero.
	Struct = struct{}

	// Interface is an empty interface.
	//
	// It says nothing; all values implement the empty interface.
	Interface = interface{}
)
