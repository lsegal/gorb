package main

/*
#include "ruby.h"
extern VALUE g_cmethod__Flip(VALUE, VALUE);

*/
import "C"
import "unsafe"
import "github.com/lsegal/gorb"
import "github.com/lsegal/gorb/test/errval"

var _ unsafe.Pointer // ignore unused import warning



//export g_cmethod__Flip
func g_cmethod__Flip(self, n uintptr) uintptr {
	go_n := int(gorb.GoInt(n))
	ret, err := errval.Flip(go_n)
  gorb.RaiseError(err)
	return gorb.IntValue(int(ret))
}



//export Init_errval
func Init_errval() {
	g_pkg := gorb.DefineModule(gorb.ModuleRoot, "Test")
	g_pkg = gorb.DefineModule(g_pkg, "Errval")

	gorb.DefineModuleFunction(g_pkg, "flip", C.g_cmethod__Flip, 1)

}

func main() { }
