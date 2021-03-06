// Package allgo contains a variety of types, functions and methods to test godoc.
package allgo

// T is an example exported type.
type T struct {
	// Name is a name.
	Name string

	// DeprecatedName is a name.
	//
	// Deprecated: Use Name instead.
	DeprecatedName string
}

// DeprecatedT is an example of a deprecated exported type.
//
// Deprecated: Use T instead.
type DeprecatedT struct{}

// F is an example exported method on T.
func (t *T) F() {}

// DeprecatedF is an example exported deprecated method on T.
//
// Deprecated: Use F instead.
func (t *T) DeprecatedF() {}

// F is an exported function.
func F() {}

// DeprecatedF is a deprecated function.
//
// Deprecated: Use F instead.
func DeprecatedF() {}

const (
	// V is a constant.
	//
	// Deprecated: Use something else.
	V = 0

	// D is another constant.
	D = 7
)

// L is a constant.
const L = 4

// B is a variable.
//
// Deprecated: Use something else.
var B = 1

var (
	// C is a var.
	C = 4

	// K is a var.
	//
	// Deprecated: Use C.
	K = 5
)
