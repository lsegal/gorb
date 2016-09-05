package gorb

import "C"
import "unsafe"

// create separate file for "//export"s as per cgo note on preambles and
// exported functions.

//export gorb_free
func gorb_free(ptr unsafe.Pointer) {
	delete(gcmap, (interface{})(ptr))
}
