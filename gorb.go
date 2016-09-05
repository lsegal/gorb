package gorb

/*
#include <stdlib.h>
#include "ruby.h"

// Ruby macro wrappers
static inline int rbmacro_NUM2INT(VALUE fix) { return NUM2INT(fix); }
static inline VALUE rbmacro_INT2NUM(int n) { return INT2NUM(n); }
static inline double rbmacro_NUM2DBL(VALUE n) { return NUM2DBL(n); }
static inline char* rbmacro_StringValueCStr(VALUE s) { return StringValueCStr(s); }

static inline VALUE rbmacro_Data_Wrap_Struct(VALUE klass, void* mark, void* free, void *ptr) {
	return Data_Wrap_Struct(klass, mark, free, ptr);
}
inline void* rbmacro_Data_Get_Struct(VALUE obj) {
	void *ret; Data_Get_Struct(obj, void, ret); return ret;
}

// extra GC helpers
extern void gorb_free(void*);
*/
import "C"
import "unsafe"

var gcmap = map[interface{}]*int{}
var revgcmap = map[*int]interface{}{}

const ModuleRoot = uintptr(0)

func GoInt(n uintptr) int {
	return int(C.rbmacro_NUM2INT(C.VALUE(n)))
}

func IntValue(i int) uintptr {
	return uintptr(C.rbmacro_INT2NUM(C.int(i)))
}

func GoBool(n uintptr) bool {
	if n == C.Qtrue {
		return true
	}
	return false
}

func BoolValue(b bool) uintptr {
	if b {
		return C.Qtrue
	}
	return C.Qfalse
}

func GoFloat(n uintptr) float64 {
	return float64(C.rbmacro_NUM2DBL(C.VALUE(n)))
}

func FloatValue(i float64) uintptr {
	return uintptr(C.rb_float_new(C.double(i)))
}

func GoString(v uintptr) string {
	return C.GoString(C.rbmacro_StringValueCStr(C.VALUE(v)))
}

func StringValue(v string) uintptr {
	str := C.CString(v)
	defer C.free(unsafe.Pointer(str))
	return uintptr(C.rb_str_new(str, C.long(len(v))))
}

func StructValue(val uintptr, obj interface{}) uintptr {
	// create a new pointer object mapped to obj so we don't pass a complex pointer
	objptr := 0
	gcmap[obj] = &objptr
	revgcmap[&objptr] = obj

	return uintptr(C.rbmacro_Data_Wrap_Struct(C.VALUE(val),
		unsafe.Pointer(nil), unsafe.Pointer(C.gorb_free), unsafe.Pointer(&objptr)))
}

func GoStruct(val uintptr) interface{} {
	if val == C.Qnil {
		return nil
	}
	objptr := (*int)(C.rbmacro_Data_Get_Struct(C.VALUE(val)))
	return revgcmap[objptr]
}

// DefineMethod defines an instance method on klass as name. The fn callback
// should be an exported C function, and arity should denote the number of
// arguments (minus the self value).
func DefineMethod(klass uintptr, name string, fn interface{}, arity int) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	C.rb_define_method(C.VALUE(klass), cname,
		(*[0]byte)(fn.(unsafe.Pointer)), C.int(arity))
}

// DefineModuleFunction defines a module function on klass as name. The fn callback
// should be an exported C function, and arity should denote the number of
// arguments (minus the self value).
func DefineModuleFunction(klass uintptr, name string, fn interface{}, arity int) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	C.rb_define_module_function(C.VALUE(klass), cname,
		(*[0]byte)(fn.(unsafe.Pointer)), C.int(arity))
}

func DefineClassMethod(klass uintptr, name string, fn interface{}, arity int) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	C.rb_define_singleton_method(C.VALUE(klass), cname,
		(*[0]byte)(fn.(unsafe.Pointer)), C.int(arity))
}

func DefineAllocator(klass uintptr, fn interface{}) {
	C.rb_define_alloc_func(C.VALUE(klass), (*[0]byte)(fn.(unsafe.Pointer)))
}

func DefineClass(parent uintptr, name string) uintptr {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	if parent == 0 {
		return uintptr(C.rb_define_class(cname, C.rb_cObject))
	}
	return uintptr(C.rb_define_class_under(C.VALUE(parent),
		cname, C.rb_cObject))
}

func DefineModule(parent uintptr, name string) uintptr {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	if parent == 0 {
		return uintptr(C.rb_define_module(cname))
	}
	return uintptr(C.rb_define_module_under(C.VALUE(parent), cname))
}
