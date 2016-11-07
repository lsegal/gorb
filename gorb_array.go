package gorb

/*
#include "ruby.h"
extern VALUE g_alloc_NativeBoolArray(VALUE);
extern VALUE g_imethod_NativeBoolArray_Each(VALUE);
extern VALUE g_imethod_NativeBoolArray_Get(VALUE, VALUE);
extern VALUE g_imethod_NativeBoolArray_Set(VALUE, VALUE, VALUE);
extern VALUE g_imethod_NativeBoolArray_Push(VALUE, VALUE);
extern VALUE g_imethod_NativeBoolArray_Size(VALUE);
extern VALUE g_imethod_NativeBoolArray_Length(VALUE);
extern VALUE g_imethod_NativeBoolArray_String(VALUE);
extern VALUE g_imethod_NativeBoolArray_Inspect(VALUE);
extern VALUE g_alloc_NativeFloatArray(VALUE);
extern VALUE g_imethod_NativeFloatArray_Each(VALUE);
extern VALUE g_imethod_NativeFloatArray_Get(VALUE, VALUE);
extern VALUE g_imethod_NativeFloatArray_Set(VALUE, VALUE, VALUE);
extern VALUE g_imethod_NativeFloatArray_Push(VALUE, VALUE);
extern VALUE g_imethod_NativeFloatArray_Size(VALUE);
extern VALUE g_imethod_NativeFloatArray_Length(VALUE);
extern VALUE g_imethod_NativeFloatArray_String(VALUE);
extern VALUE g_imethod_NativeFloatArray_Inspect(VALUE);
extern VALUE g_alloc_NativeIntArray(VALUE);
extern VALUE g_imethod_NativeIntArray_Each(VALUE);
extern VALUE g_imethod_NativeIntArray_Get(VALUE, VALUE);
extern VALUE g_imethod_NativeIntArray_Set(VALUE, VALUE, VALUE);
extern VALUE g_imethod_NativeIntArray_Push(VALUE, VALUE);
extern VALUE g_imethod_NativeIntArray_Size(VALUE);
extern VALUE g_imethod_NativeIntArray_Length(VALUE);
extern VALUE g_imethod_NativeIntArray_String(VALUE);
extern VALUE g_imethod_NativeIntArray_Inspect(VALUE);
extern VALUE g_alloc_NativeStringArray(VALUE);
extern VALUE g_imethod_NativeStringArray_Each(VALUE);
extern VALUE g_imethod_NativeStringArray_Get(VALUE, VALUE);
extern VALUE g_imethod_NativeStringArray_Set(VALUE, VALUE, VALUE);
extern VALUE g_imethod_NativeStringArray_Push(VALUE, VALUE);
extern VALUE g_imethod_NativeStringArray_Size(VALUE);
extern VALUE g_imethod_NativeStringArray_Length(VALUE);
extern VALUE g_imethod_NativeStringArray_String(VALUE);
extern VALUE g_imethod_NativeStringArray_Inspect(VALUE);

*/
import "C"
import "unsafe"
import "github.com/lsegal/gorb/native"

var _ unsafe.Pointer // ignore unused import warning

var g_class_NativeBoolArray uintptr
var g_class_NativeFloatArray uintptr
var g_class_NativeIntArray uintptr
var g_class_NativeStringArray uintptr

func g_val2ptr_NativeBoolArray(obj uintptr) *native.NativeBoolArray {
	return (*native.NativeBoolArray)(GoStruct(obj))
}

//export g_alloc_NativeBoolArray
func g_alloc_NativeBoolArray(klass uintptr) uintptr {
	return g_classinit_NativeBoolArray(klass, &native.NativeBoolArray{})
}

func g_classinit_NativeBoolArray(klass uintptr, obj *native.NativeBoolArray) uintptr {
	return StructValue(klass, unsafe.Pointer(obj))
}

//export g_imethod_NativeBoolArray_Each
func g_imethod_NativeBoolArray_Each(self uintptr) uintptr {
	if e := EnumFor(self, StringValue("each")); e != C.Qnil {
		return e
	}
	go_obj := g_val2ptr_NativeBoolArray(self)
	go_obj.Each(block__g_imethod_NativeBoolArray_Each)
	return C.Qnil
}

func block__g_imethod_NativeBoolArray_Each(arg0 bool) {
	rb_arg0 := BoolValue(bool(arg0))
	Yield(rb_arg0)
}

//export g_imethod_NativeBoolArray_Get
func g_imethod_NativeBoolArray_Get(self, idx uintptr) uintptr {
	go_obj := g_val2ptr_NativeBoolArray(self)
	go_idx := int(GoInt(idx))
	ret := go_obj.Get(go_idx)
	return BoolValue(bool(ret))
}

//export g_imethod_NativeBoolArray_Set
func g_imethod_NativeBoolArray_Set(self, idx, value uintptr) uintptr {
	go_obj := g_val2ptr_NativeBoolArray(self)
	go_idx := int(GoInt(idx))
	go_value := bool(GoBool(value))
	go_obj.Set(go_idx, go_value)
	return C.Qnil
}

//export g_imethod_NativeBoolArray_Push
func g_imethod_NativeBoolArray_Push(self, value uintptr) uintptr {
	go_obj := g_val2ptr_NativeBoolArray(self)
	go_value := bool(GoBool(value))
	go_obj.Push(go_value)
	return C.Qnil
}

//export g_imethod_NativeBoolArray_Size
func g_imethod_NativeBoolArray_Size(self uintptr) uintptr {
	go_obj := g_val2ptr_NativeBoolArray(self)
	ret := go_obj.Size()
	return IntValue(int(ret))
}

//export g_imethod_NativeBoolArray_Length
func g_imethod_NativeBoolArray_Length(self uintptr) uintptr {
	go_obj := g_val2ptr_NativeBoolArray(self)
	ret := go_obj.Length()
	return IntValue(int(ret))
}

//export g_imethod_NativeBoolArray_String
func g_imethod_NativeBoolArray_String(self uintptr) uintptr {
	go_obj := g_val2ptr_NativeBoolArray(self)
	ret := go_obj.String()
	return StringValue(string(ret))
}

//export g_imethod_NativeBoolArray_Inspect
func g_imethod_NativeBoolArray_Inspect(self uintptr) uintptr {
	go_obj := g_val2ptr_NativeBoolArray(self)
	ret := go_obj.Inspect()
	return StringValue(string(ret))
}

func g_val2ptr_NativeFloatArray(obj uintptr) *native.NativeFloatArray {
	return (*native.NativeFloatArray)(GoStruct(obj))
}

//export g_alloc_NativeFloatArray
func g_alloc_NativeFloatArray(klass uintptr) uintptr {
	return g_classinit_NativeFloatArray(klass, &native.NativeFloatArray{})
}

func g_classinit_NativeFloatArray(klass uintptr, obj *native.NativeFloatArray) uintptr {
	return StructValue(klass, unsafe.Pointer(obj))
}

//export g_imethod_NativeFloatArray_Each
func g_imethod_NativeFloatArray_Each(self uintptr) uintptr {
	if e := EnumFor(self, StringValue("each")); e != C.Qnil {
		return e
	}
	go_obj := g_val2ptr_NativeFloatArray(self)
	go_obj.Each(block__g_imethod_NativeFloatArray_Each)
	return C.Qnil
}

func block__g_imethod_NativeFloatArray_Each(arg0 float64) {
	rb_arg0 := FloatValue(float64(arg0))
	Yield(rb_arg0)
}

//export g_imethod_NativeFloatArray_Get
func g_imethod_NativeFloatArray_Get(self, idx uintptr) uintptr {
	go_obj := g_val2ptr_NativeFloatArray(self)
	go_idx := int(GoInt(idx))
	ret := go_obj.Get(go_idx)
	return FloatValue(float64(ret))
}

//export g_imethod_NativeFloatArray_Set
func g_imethod_NativeFloatArray_Set(self, idx, value uintptr) uintptr {
	go_obj := g_val2ptr_NativeFloatArray(self)
	go_idx := int(GoInt(idx))
	go_value := float64(GoFloat(value))
	go_obj.Set(go_idx, go_value)
	return C.Qnil
}

//export g_imethod_NativeFloatArray_Push
func g_imethod_NativeFloatArray_Push(self, value uintptr) uintptr {
	go_obj := g_val2ptr_NativeFloatArray(self)
	go_value := float64(GoFloat(value))
	go_obj.Push(go_value)
	return C.Qnil
}

//export g_imethod_NativeFloatArray_Size
func g_imethod_NativeFloatArray_Size(self uintptr) uintptr {
	go_obj := g_val2ptr_NativeFloatArray(self)
	ret := go_obj.Size()
	return IntValue(int(ret))
}

//export g_imethod_NativeFloatArray_Length
func g_imethod_NativeFloatArray_Length(self uintptr) uintptr {
	go_obj := g_val2ptr_NativeFloatArray(self)
	ret := go_obj.Length()
	return IntValue(int(ret))
}

//export g_imethod_NativeFloatArray_String
func g_imethod_NativeFloatArray_String(self uintptr) uintptr {
	go_obj := g_val2ptr_NativeFloatArray(self)
	ret := go_obj.String()
	return StringValue(string(ret))
}

//export g_imethod_NativeFloatArray_Inspect
func g_imethod_NativeFloatArray_Inspect(self uintptr) uintptr {
	go_obj := g_val2ptr_NativeFloatArray(self)
	ret := go_obj.Inspect()
	return StringValue(string(ret))
}

func g_val2ptr_NativeIntArray(obj uintptr) *native.NativeIntArray {
	return (*native.NativeIntArray)(GoStruct(obj))
}

//export g_alloc_NativeIntArray
func g_alloc_NativeIntArray(klass uintptr) uintptr {
	return g_classinit_NativeIntArray(klass, &native.NativeIntArray{})
}

func g_classinit_NativeIntArray(klass uintptr, obj *native.NativeIntArray) uintptr {
	return StructValue(klass, unsafe.Pointer(obj))
}

//export g_imethod_NativeIntArray_Each
func g_imethod_NativeIntArray_Each(self uintptr) uintptr {
	if e := EnumFor(self, StringValue("each")); e != C.Qnil {
		return e
	}
	go_obj := g_val2ptr_NativeIntArray(self)
	go_obj.Each(block__g_imethod_NativeIntArray_Each)
	return C.Qnil
}

func block__g_imethod_NativeIntArray_Each(arg0 int) {
	rb_arg0 := IntValue(int(arg0))
	Yield(rb_arg0)
}

//export g_imethod_NativeIntArray_Get
func g_imethod_NativeIntArray_Get(self, idx uintptr) uintptr {
	go_obj := g_val2ptr_NativeIntArray(self)
	go_idx := int(GoInt(idx))
	ret := go_obj.Get(go_idx)
	return IntValue(int(ret))
}

//export g_imethod_NativeIntArray_Set
func g_imethod_NativeIntArray_Set(self, idx, value uintptr) uintptr {
	go_obj := g_val2ptr_NativeIntArray(self)
	go_idx := int(GoInt(idx))
	go_value := int(GoInt(value))
	go_obj.Set(go_idx, go_value)
	return C.Qnil
}

//export g_imethod_NativeIntArray_Push
func g_imethod_NativeIntArray_Push(self, value uintptr) uintptr {
	go_obj := g_val2ptr_NativeIntArray(self)
	go_value := int(GoInt(value))
	go_obj.Push(go_value)
	return C.Qnil
}

//export g_imethod_NativeIntArray_Size
func g_imethod_NativeIntArray_Size(self uintptr) uintptr {
	go_obj := g_val2ptr_NativeIntArray(self)
	ret := go_obj.Size()
	return IntValue(int(ret))
}

//export g_imethod_NativeIntArray_Length
func g_imethod_NativeIntArray_Length(self uintptr) uintptr {
	go_obj := g_val2ptr_NativeIntArray(self)
	ret := go_obj.Length()
	return IntValue(int(ret))
}

//export g_imethod_NativeIntArray_String
func g_imethod_NativeIntArray_String(self uintptr) uintptr {
	go_obj := g_val2ptr_NativeIntArray(self)
	ret := go_obj.String()
	return StringValue(string(ret))
}

//export g_imethod_NativeIntArray_Inspect
func g_imethod_NativeIntArray_Inspect(self uintptr) uintptr {
	go_obj := g_val2ptr_NativeIntArray(self)
	ret := go_obj.Inspect()
	return StringValue(string(ret))
}

func g_val2ptr_NativeStringArray(obj uintptr) *native.NativeStringArray {
	return (*native.NativeStringArray)(GoStruct(obj))
}

//export g_alloc_NativeStringArray
func g_alloc_NativeStringArray(klass uintptr) uintptr {
	return g_classinit_NativeStringArray(klass, &native.NativeStringArray{})
}

func g_classinit_NativeStringArray(klass uintptr, obj *native.NativeStringArray) uintptr {
	return StructValue(klass, unsafe.Pointer(obj))
}

//export g_imethod_NativeStringArray_Each
func g_imethod_NativeStringArray_Each(self uintptr) uintptr {
	if e := EnumFor(self, StringValue("each")); e != C.Qnil {
		return e
	}
	go_obj := g_val2ptr_NativeStringArray(self)
	go_obj.Each(block__g_imethod_NativeStringArray_Each)
	return C.Qnil
}

func block__g_imethod_NativeStringArray_Each(arg0 string) {
	rb_arg0 := StringValue(string(arg0))
	Yield(rb_arg0)
}

//export g_imethod_NativeStringArray_Get
func g_imethod_NativeStringArray_Get(self, idx uintptr) uintptr {
	go_obj := g_val2ptr_NativeStringArray(self)
	go_idx := int(GoInt(idx))
	ret := go_obj.Get(go_idx)
	return StringValue(string(ret))
}

//export g_imethod_NativeStringArray_Set
func g_imethod_NativeStringArray_Set(self, idx, value uintptr) uintptr {
	go_obj := g_val2ptr_NativeStringArray(self)
	go_idx := int(GoInt(idx))
	go_value := string(GoString(value))
	go_obj.Set(go_idx, go_value)
	return C.Qnil
}

//export g_imethod_NativeStringArray_Push
func g_imethod_NativeStringArray_Push(self, value uintptr) uintptr {
	go_obj := g_val2ptr_NativeStringArray(self)
	go_value := string(GoString(value))
	go_obj.Push(go_value)
	return C.Qnil
}

//export g_imethod_NativeStringArray_Size
func g_imethod_NativeStringArray_Size(self uintptr) uintptr {
	go_obj := g_val2ptr_NativeStringArray(self)
	ret := go_obj.Size()
	return IntValue(int(ret))
}

//export g_imethod_NativeStringArray_Length
func g_imethod_NativeStringArray_Length(self uintptr) uintptr {
	go_obj := g_val2ptr_NativeStringArray(self)
	ret := go_obj.Length()
	return IntValue(int(ret))
}

//export g_imethod_NativeStringArray_String
func g_imethod_NativeStringArray_String(self uintptr) uintptr {
	go_obj := g_val2ptr_NativeStringArray(self)
	ret := go_obj.String()
	return StringValue(string(ret))
}

//export g_imethod_NativeStringArray_Inspect
func g_imethod_NativeStringArray_Inspect(self uintptr) uintptr {
	go_obj := g_val2ptr_NativeStringArray(self)
	ret := go_obj.Inspect()
	return StringValue(string(ret))
}

func init() {
	g_pkg := DefineModule(ModuleRoot, "Gorb")

	g_class_NativeBoolArray = DefineClass(g_pkg, "NativeBoolArray")
	DefineAllocator(g_class_NativeBoolArray, C.g_alloc_NativeBoolArray)
	DefineMethod(g_class_NativeBoolArray, "each", C.g_imethod_NativeBoolArray_Each, 0)
	DefineMethod(g_class_NativeBoolArray, "[]", C.g_imethod_NativeBoolArray_Get, 1)
	DefineMethod(g_class_NativeBoolArray, "[]=", C.g_imethod_NativeBoolArray_Set, 2)
	DefineMethod(g_class_NativeBoolArray, "push", C.g_imethod_NativeBoolArray_Push, 1)
	DefineMethod(g_class_NativeBoolArray, "size", C.g_imethod_NativeBoolArray_Size, 0)
	DefineMethod(g_class_NativeBoolArray, "length", C.g_imethod_NativeBoolArray_Length, 0)
	DefineMethod(g_class_NativeBoolArray, "to_s", C.g_imethod_NativeBoolArray_String, 0)
	DefineMethod(g_class_NativeBoolArray, "inspect", C.g_imethod_NativeBoolArray_Inspect, 0)
	g_class_NativeFloatArray = DefineClass(g_pkg, "NativeFloatArray")
	DefineAllocator(g_class_NativeFloatArray, C.g_alloc_NativeFloatArray)
	DefineMethod(g_class_NativeFloatArray, "each", C.g_imethod_NativeFloatArray_Each, 0)
	DefineMethod(g_class_NativeFloatArray, "[]", C.g_imethod_NativeFloatArray_Get, 1)
	DefineMethod(g_class_NativeFloatArray, "[]=", C.g_imethod_NativeFloatArray_Set, 2)
	DefineMethod(g_class_NativeFloatArray, "push", C.g_imethod_NativeFloatArray_Push, 1)
	DefineMethod(g_class_NativeFloatArray, "size", C.g_imethod_NativeFloatArray_Size, 0)
	DefineMethod(g_class_NativeFloatArray, "length", C.g_imethod_NativeFloatArray_Length, 0)
	DefineMethod(g_class_NativeFloatArray, "to_s", C.g_imethod_NativeFloatArray_String, 0)
	DefineMethod(g_class_NativeFloatArray, "inspect", C.g_imethod_NativeFloatArray_Inspect, 0)
	g_class_NativeIntArray = DefineClass(g_pkg, "NativeIntArray")
	DefineAllocator(g_class_NativeIntArray, C.g_alloc_NativeIntArray)
	DefineMethod(g_class_NativeIntArray, "each", C.g_imethod_NativeIntArray_Each, 0)
	DefineMethod(g_class_NativeIntArray, "[]", C.g_imethod_NativeIntArray_Get, 1)
	DefineMethod(g_class_NativeIntArray, "[]=", C.g_imethod_NativeIntArray_Set, 2)
	DefineMethod(g_class_NativeIntArray, "push", C.g_imethod_NativeIntArray_Push, 1)
	DefineMethod(g_class_NativeIntArray, "size", C.g_imethod_NativeIntArray_Size, 0)
	DefineMethod(g_class_NativeIntArray, "length", C.g_imethod_NativeIntArray_Length, 0)
	DefineMethod(g_class_NativeIntArray, "to_s", C.g_imethod_NativeIntArray_String, 0)
	DefineMethod(g_class_NativeIntArray, "inspect", C.g_imethod_NativeIntArray_Inspect, 0)
	g_class_NativeStringArray = DefineClass(g_pkg, "NativeStringArray")
	DefineAllocator(g_class_NativeStringArray, C.g_alloc_NativeStringArray)
	DefineMethod(g_class_NativeStringArray, "each", C.g_imethod_NativeStringArray_Each, 0)
	DefineMethod(g_class_NativeStringArray, "[]", C.g_imethod_NativeStringArray_Get, 1)
	DefineMethod(g_class_NativeStringArray, "[]=", C.g_imethod_NativeStringArray_Set, 2)
	DefineMethod(g_class_NativeStringArray, "push", C.g_imethod_NativeStringArray_Push, 1)
	DefineMethod(g_class_NativeStringArray, "size", C.g_imethod_NativeStringArray_Size, 0)
	DefineMethod(g_class_NativeStringArray, "length", C.g_imethod_NativeStringArray_Length, 0)
	DefineMethod(g_class_NativeStringArray, "to_s", C.g_imethod_NativeStringArray_String, 0)
	DefineMethod(g_class_NativeStringArray, "inspect", C.g_imethod_NativeStringArray_Inspect, 0)

}
