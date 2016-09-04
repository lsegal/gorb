package gorb

import "C"
import "unsafe"

// create separate file for "//export"s as per cgo note on preambles and
// exported functions.

//export gorb_free
func gorb_free(ptr unsafe.Pointer) {
	objif := (interface{})(ptr)
	objptr := gcmap[objif]
	delete(gcmap, objif)
	delete(revgcmap, objptr)
}
