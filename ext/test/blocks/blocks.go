
package main

/*
#include "ruby.h"
extern VALUE g_cmethod__Double(VALUE, VALUE);

*/
import "C"
import "unsafe"
import "github.com/lsegal/gorb"
import "github.com/lsegal/gorb/test/blocks"

var _ unsafe.Pointer // ignore unused import warning



//export g_cmethod__Double
func g_cmethod__Double(self, val uintptr) uintptr {
	go_val := int(gorb.GoInt(val))
	ret := blocks.Double(go_val, block__g_cmethod__Double)
	return gorb.IntValue(int(ret))
}

func block__g_cmethod__Double(arg0 int) (int) {
	rb_arg0 := gorb.IntValue(int(arg0))
	ret := gorb.Yield(rb_arg0)
	return int(gorb.GoInt(ret))
}



//export Init_blocks
func Init_blocks() {
	g_pkg := gorb.DefineModule(gorb.ModuleRoot, "Test")
	g_pkg = gorb.DefineModule(g_pkg, "Blocks")

	gorb.DefineModuleFunction(g_pkg, "double", C.g_cmethod__Double, 1)

}

func main() {}
