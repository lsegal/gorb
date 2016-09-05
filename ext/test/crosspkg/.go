
package main

/*
#include "ruby.h"

*/
import "C"
import "unsafe"
import "github.com/lsegal/gorb"


var _ unsafe.Pointer // ignore unused import warning




//export Init_
func Init_() {
	g_pkg := gorb.DefineModule(gorb.ModuleRoot, "Test")
	g_pkg = gorb.DefineModule(g_pkg, "Crosspkg")


}

func main() {}
