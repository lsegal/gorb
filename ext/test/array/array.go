
package main

/*
#include "ruby.h"
extern VALUE g_cmethod__ReverseArray(VALUE, VALUE);

*/
import "C"
import "unsafe"
import "github.com/lsegal/gorb"
import "github.com/lsegal/gorb/test/array"

var _ unsafe.Pointer // ignore unused import warning



//export g_cmethod__ReverseArray
func g_cmethod__ReverseArray(self, list uintptr) uintptr {
	go_list := []string(gorb.GoStringArray(list))
	ret := array.ReverseArray(go_list)
	return gorb.ArrayStringValue([]string(ret))
}



//export Init_array
func Init_array() {
	g_pkg := gorb.DefineModule(gorb.ModuleRoot, "Test")
	g_pkg = gorb.DefineModule(g_pkg, "Array")

	gorb.DefineModuleFunction(g_pkg, "reverse_array", C.g_cmethod__ReverseArray, 1)

}

func main() {}
