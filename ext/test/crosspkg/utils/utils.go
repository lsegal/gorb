
package main

/*
#include "ruby.h"
extern VALUE g_cmethod__ToHSV(VALUE, VALUE);

*/
import "C"
import "unsafe"
import "github.com/lsegal/gorb"
import "github.com/lsegal/gorb/test/crosspkg/data"
import "github.com/lsegal/gorb/test/crosspkg/utils"

var _ unsafe.Pointer // ignore unused import warning



//export g_cmethod__ToHSV
func g_cmethod__ToHSV(self, rgb uintptr) uintptr {
	go_rgb := (*data.RGB)(gorb.GoStruct(rgb))
	ret := utils.ToHSV(*go_rgb)
	return gorb.StructValue(gorb.ObjAtPath("Test::Crosspkg::Data::HSV"), unsafe.Pointer(&ret))
}



//export Init_utils
func Init_utils() {
	g_pkg := gorb.DefineModule(gorb.ModuleRoot, "Test")
	g_pkg = gorb.DefineModule(g_pkg, "Crosspkg")
	g_pkg = gorb.DefineModule(g_pkg, "Utils")

	gorb.DefineModuleFunction(g_pkg, "to_hsv", C.g_cmethod__ToHSV, 1)

}

func main() {}
