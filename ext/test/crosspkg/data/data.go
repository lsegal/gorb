
package main

/*
#include "ruby.h"
extern VALUE g_alloc_RGB(VALUE);
extern VALUE g_imethod_RGB_R(VALUE);
extern VALUE g_imethod_RGB_R__set(VALUE, VALUE);
extern VALUE g_imethod_RGB_G(VALUE);
extern VALUE g_imethod_RGB_G__set(VALUE, VALUE);
extern VALUE g_imethod_RGB_B(VALUE);
extern VALUE g_imethod_RGB_B__set(VALUE, VALUE);
extern VALUE g_alloc_HSV(VALUE);
extern VALUE g_imethod_HSV_H(VALUE);
extern VALUE g_imethod_HSV_H__set(VALUE, VALUE);
extern VALUE g_imethod_HSV_S(VALUE);
extern VALUE g_imethod_HSV_S__set(VALUE, VALUE);
extern VALUE g_imethod_HSV_V(VALUE);
extern VALUE g_imethod_HSV_V__set(VALUE, VALUE);
extern VALUE g_imethod_HSV_String(VALUE);
extern VALUE g_imethod_HSV_Inspect(VALUE);

*/
import "C"
import "unsafe"
import "github.com/lsegal/gorb"
import "github.com/lsegal/gorb/test/crosspkg/data"

var _ unsafe.Pointer // ignore unused import warning

var g_class_RGB uintptr
var g_class_HSV uintptr


func g_val2ptr_RGB(obj uintptr) *data.RGB {
	return (*data.RGB)(gorb.GoStruct(obj))
}

//export g_alloc_RGB
func g_alloc_RGB(klass uintptr) uintptr {
	return g_classinit_RGB(klass, &data.RGB{})
}

func g_classinit_RGB(klass uintptr, obj *data.RGB) uintptr {
	return gorb.StructValue(klass, unsafe.Pointer(obj))
}


//export g_imethod_RGB_R
func g_imethod_RGB_R(self uintptr) uintptr {
	obj := g_val2ptr_RGB(self)
	return gorb.IntValue(int(obj.R))
}

//export g_imethod_RGB_R__set
func g_imethod_RGB_R__set(self, val uintptr) uintptr {
	obj := g_val2ptr_RGB(self)
	obj.R = int(gorb.GoInt(val))
	return val
}


//export g_imethod_RGB_G
func g_imethod_RGB_G(self uintptr) uintptr {
	obj := g_val2ptr_RGB(self)
	return gorb.IntValue(int(obj.G))
}

//export g_imethod_RGB_G__set
func g_imethod_RGB_G__set(self, val uintptr) uintptr {
	obj := g_val2ptr_RGB(self)
	obj.G = int(gorb.GoInt(val))
	return val
}


//export g_imethod_RGB_B
func g_imethod_RGB_B(self uintptr) uintptr {
	obj := g_val2ptr_RGB(self)
	return gorb.IntValue(int(obj.B))
}

//export g_imethod_RGB_B__set
func g_imethod_RGB_B__set(self, val uintptr) uintptr {
	obj := g_val2ptr_RGB(self)
	obj.B = int(gorb.GoInt(val))
	return val
}


func g_val2ptr_HSV(obj uintptr) *data.HSV {
	return (*data.HSV)(gorb.GoStruct(obj))
}

//export g_alloc_HSV
func g_alloc_HSV(klass uintptr) uintptr {
	return g_classinit_HSV(klass, &data.HSV{})
}

func g_classinit_HSV(klass uintptr, obj *data.HSV) uintptr {
	return gorb.StructValue(klass, unsafe.Pointer(obj))
}


//export g_imethod_HSV_H
func g_imethod_HSV_H(self uintptr) uintptr {
	obj := g_val2ptr_HSV(self)
	return gorb.FloatValue(float64(obj.H))
}

//export g_imethod_HSV_H__set
func g_imethod_HSV_H__set(self, val uintptr) uintptr {
	obj := g_val2ptr_HSV(self)
	obj.H = float64(gorb.GoFloat(val))
	return val
}


//export g_imethod_HSV_S
func g_imethod_HSV_S(self uintptr) uintptr {
	obj := g_val2ptr_HSV(self)
	return gorb.FloatValue(float64(obj.S))
}

//export g_imethod_HSV_S__set
func g_imethod_HSV_S__set(self, val uintptr) uintptr {
	obj := g_val2ptr_HSV(self)
	obj.S = float64(gorb.GoFloat(val))
	return val
}


//export g_imethod_HSV_V
func g_imethod_HSV_V(self uintptr) uintptr {
	obj := g_val2ptr_HSV(self)
	return gorb.FloatValue(float64(obj.V))
}

//export g_imethod_HSV_V__set
func g_imethod_HSV_V__set(self, val uintptr) uintptr {
	obj := g_val2ptr_HSV(self)
	obj.V = float64(gorb.GoFloat(val))
	return val
}


//export g_imethod_HSV_String
func g_imethod_HSV_String(self uintptr) uintptr {
	go_obj := g_val2ptr_HSV(self)
	ret := go_obj.String()
	return gorb.StringValue(string(ret))
}


//export g_imethod_HSV_Inspect
func g_imethod_HSV_Inspect(self uintptr) uintptr {
	go_obj := g_val2ptr_HSV(self)
	ret := go_obj.Inspect()
	return gorb.StringValue(string(ret))
}



//export Init_data
func Init_data() {
	g_pkg := gorb.DefineModule(gorb.ModuleRoot, "Test")
	g_pkg = gorb.DefineModule(g_pkg, "Crosspkg")
	g_pkg = gorb.DefineModule(g_pkg, "Data")

	g_class_RGB = gorb.DefineClass(g_pkg, "RGB")
	gorb.DefineAllocator(g_class_RGB, C.g_alloc_RGB)
	gorb.DefineMethod(g_class_RGB, "r", C.g_imethod_RGB_R, 0)
	gorb.DefineMethod(g_class_RGB, "r=", C.g_imethod_RGB_R__set, 1)
	gorb.DefineMethod(g_class_RGB, "g", C.g_imethod_RGB_G, 0)
	gorb.DefineMethod(g_class_RGB, "g=", C.g_imethod_RGB_G__set, 1)
	gorb.DefineMethod(g_class_RGB, "b", C.g_imethod_RGB_B, 0)
	gorb.DefineMethod(g_class_RGB, "b=", C.g_imethod_RGB_B__set, 1)
	g_class_HSV = gorb.DefineClass(g_pkg, "HSV")
	gorb.DefineAllocator(g_class_HSV, C.g_alloc_HSV)
	gorb.DefineMethod(g_class_HSV, "h", C.g_imethod_HSV_H, 0)
	gorb.DefineMethod(g_class_HSV, "h=", C.g_imethod_HSV_H__set, 1)
	gorb.DefineMethod(g_class_HSV, "s", C.g_imethod_HSV_S, 0)
	gorb.DefineMethod(g_class_HSV, "s=", C.g_imethod_HSV_S__set, 1)
	gorb.DefineMethod(g_class_HSV, "v", C.g_imethod_HSV_V, 0)
	gorb.DefineMethod(g_class_HSV, "v=", C.g_imethod_HSV_V__set, 1)
	gorb.DefineMethod(g_class_HSV, "to_s", C.g_imethod_HSV_String, 0)
	gorb.DefineMethod(g_class_HSV, "inspect", C.g_imethod_HSV_Inspect, 0)

}

func main() {}
