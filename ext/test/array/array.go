package main

/*
#include "ruby.h"
extern VALUE g_cmethod__ReverseArray(VALUE, VALUE);
extern VALUE g_cmethod__MutateArray(VALUE, VALUE);
extern VALUE g_cmethod__MutateIntArray(VALUE, VALUE);

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
	return gorb.ArrayStringValue((ret))
}


//export g_cmethod__MutateArray
func g_cmethod__MutateArray(self, list uintptr) uintptr {
	if e := gorb.EnumFor(self, gorb.StringValue("mutate_array"), list); e != C.Qnil {
		return e
	}
	go_list := []string(gorb.GoStringArray(list))
	array.MutateArray(go_list, block__g_cmethod__MutateArray)
	return C.Qnil
}

func block__g_cmethod__MutateArray(arg0 *[]string) () {
	rb_arg0 := gorb.ArrayStringValue((arg0))
	gorb.Yield(rb_arg0)
}


//export g_cmethod__MutateIntArray
func g_cmethod__MutateIntArray(self, list uintptr) uintptr {
	if e := gorb.EnumFor(self, gorb.StringValue("mutate_int_array"), list); e != C.Qnil {
		return e
	}
	go_list := []int(gorb.GoIntArray(list))
	array.MutateIntArray(go_list, block__g_cmethod__MutateIntArray)
	return C.Qnil
}

func block__g_cmethod__MutateIntArray(arg0 *[]int) () {
	rb_arg0 := gorb.ArrayIntValue((arg0))
	gorb.Yield(rb_arg0)
}



//export Init_array
func Init_array() {
	g_pkg := gorb.DefineModule(gorb.ModuleRoot, "Test")
	g_pkg = gorb.DefineModule(g_pkg, "Array")

	gorb.DefineModuleFunction(g_pkg, "reverse_array", C.g_cmethod__ReverseArray, 1)
	gorb.DefineModuleFunction(g_pkg, "mutate_array", C.g_cmethod__MutateArray, 1)
	gorb.DefineModuleFunction(g_pkg, "mutate_int_array", C.g_cmethod__MutateIntArray, 1)

}

func main() { }
