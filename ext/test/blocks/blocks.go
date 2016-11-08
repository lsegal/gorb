package main

/*
#include "ruby.h"
extern VALUE g_cmethod__DoWith(VALUE, VALUE);

*/
import "C"
import "unsafe"
import "github.com/lsegal/gorb"
import "github.com/lsegal/gorb/test/blocks"

var _ unsafe.Pointer // ignore unused import warning



//export g_cmethod__DoWith
func g_cmethod__DoWith(self, val uintptr) uintptr {
	if e := gorb.EnumFor(self, gorb.StringValue("do_with"), val); e != C.Qnil {
		return e
	}
	go_val := int(gorb.GoInt(val))
	ret := blocks.DoWith(go_val, block__g_cmethod__DoWith)
	return gorb.IntValue(int(ret))
}

func block__g_cmethod__DoWith(arg0 int) (int) {
	rb_arg0 := gorb.IntValue(int(arg0))
	ret := gorb.Yield(rb_arg0)
	return int(gorb.GoInt(ret))
}



//export Init_blocks
func Init_blocks() {
	g_pkg := gorb.DefineModule(gorb.ModuleRoot, "Test")
	g_pkg = gorb.DefineModule(g_pkg, "Blocks")

	gorb.DefineModuleFunction(g_pkg, "do_with", C.g_cmethod__DoWith, 1)

}

func main() { }
