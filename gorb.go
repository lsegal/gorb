package gorb

/*
#include <stdlib.h>
#include "ruby.h"

// Ruby macro wrappers
static inline int rbmacro_NUM2INT(VALUE fix) { return NUM2INT(fix); }
static inline VALUE rbmacro_INT2NUM(int n) { return INT2NUM(n); }
static inline double rbmacro_NUM2DBL(VALUE n) { return NUM2DBL(n); }
static inline char* rbmacro_StringValueCStr(VALUE s) { return StringValueCStr(s); }
static inline VALUE rbmacro_Data_Wrap_Struct(VALUE klass, void* mark, void* free, char *ptr) {
	return Data_Wrap_Struct(klass, mark, free, (void*)ptr);
}
static inline void* rbmacro_Data_Get_Struct(VALUE obj) {
	void *ret; Data_Get_Struct(obj, void, ret); return ret;
}
static inline void rbmacro_ary_set(VALUE obj, int idx, VALUE val) {
	RARRAY_ASET(obj, idx, val);
}
static inline VALUE rbmacro_ary_get(VALUE obj, int idx) {
	return RARRAY_AREF(obj, idx);
}
static inline long rbmacro_ary_len(VALUE obj) {
	return RARRAY_LEN(obj);
}

// extra GC helpers
extern void gorb_free(void*);
*/
import "C"
import (
	"unsafe"

	"github.com/lsegal/gorb/native"
)

var gcmap = map[interface{}]bool{}
var idProcCall C.ID
var idEnumFor C.ID

func init() {
	call := C.CString("call")
	defer C.free(unsafe.Pointer(call))
	enumFor := C.CString("enum_for")
	defer C.free(unsafe.Pointer(enumFor))
	idProcCall = C.rb_intern(call)
	idEnumFor = C.rb_intern(enumFor)
}

const ModuleRoot = uintptr(0)

func GoStringArray(arr uintptr) []string {
	list := make([]string, int(C.rbmacro_ary_len(C.VALUE(arr))))
	for i := 0; i < int(C.rbmacro_ary_len(C.VALUE(arr))); i++ {
		list[i] = GoString(uintptr(C.rbmacro_ary_get(C.VALUE(arr), C.int(i))))
	}
	return list
}

func ArrayStringValue(arr *[]string) uintptr {
	return g_classinit_NativeStringArray(g_class_NativeStringArray,
		&native.NativeStringArray{List: arr})
}

func GoIntArray(arr uintptr) []int {
	list := make([]int, int(C.rbmacro_ary_len(C.VALUE(arr))))
	for i := 0; i < int(C.rbmacro_ary_len(C.VALUE(arr))); i++ {
		list[i] = GoInt(uintptr(C.rbmacro_ary_get(C.VALUE(arr), C.int(i))))
	}
	return list
}

func ArrayIntValue(arr *[]int) uintptr {
	return g_classinit_NativeIntArray(g_class_NativeIntArray,
		&native.NativeIntArray{List: arr})
}

func GoBoolArray(arr uintptr) []bool {
	list := make([]bool, int(C.rbmacro_ary_len(C.VALUE(arr))))
	for i := 0; i < int(C.rbmacro_ary_len(C.VALUE(arr))); i++ {
		list[i] = GoBool(uintptr(C.rbmacro_ary_get(C.VALUE(arr), C.int(i))))
	}
	return list
}

func ArrayBoolValue(arr *[]bool) uintptr {
	return g_classinit_NativeBoolArray(g_class_NativeBoolArray,
		&native.NativeBoolArray{List: arr})
}

func GoFloatArray(arr uintptr) []float64 {
	list := make([]float64, int(C.rbmacro_ary_len(C.VALUE(arr))))
	for i := 0; i < int(C.rbmacro_ary_len(C.VALUE(arr))); i++ {
		list[i] = GoFloat(uintptr(C.rbmacro_ary_get(C.VALUE(arr), C.int(i))))
	}
	return list
}

func ArrayFloatValue(arr *[]float64) uintptr {
	return g_classinit_NativeFloatArray(g_class_NativeFloatArray,
		&native.NativeFloatArray{List: arr})
}

func GoStructArray(arr uintptr) []unsafe.Pointer {
	list := make([]unsafe.Pointer, int(C.rbmacro_ary_len(C.VALUE(arr))))
	for i := 0; i < int(C.rbmacro_ary_len(C.VALUE(arr))); i++ {
		list[i] = GoStruct(uintptr(C.rbmacro_ary_get(C.VALUE(arr), C.int(i))))
	}
	return list
}

func ArrayStructValue(typ uintptr, arr []unsafe.Pointer) uintptr {
	list := C.rb_ary_new2(C.long(len(arr)))
	for _, v := range arr {
		C.rb_ary_push(list, C.VALUE(StructValue(typ, v)))
	}
	return uintptr(list)
}

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

func StructValue(val uintptr, obj unsafe.Pointer) uintptr {
	if obj == nil {
		return C.Qnil
	}
	gcmap[obj] = true
	return uintptr(C.rbmacro_Data_Wrap_Struct(C.VALUE(val),
		unsafe.Pointer(nil), unsafe.Pointer(C.gorb_free), (*C.char)(obj)))
}

func GoStruct(val uintptr) unsafe.Pointer {
	if val == C.Qnil {
		return nil
	}
	return C.rbmacro_Data_Get_Struct(C.VALUE(val))
}

func ObjAtPath(path string) uintptr {
	cpath := C.CString(path)
	defer C.free(unsafe.Pointer(cpath))
	return uintptr(C.rb_path2class(cpath))
}

func Yield(values ...uintptr) uintptr {
	if len(values) == 0 {
		return uintptr(C.rb_yield(C.Qundef))
	}
	return uintptr(C.rb_yield_values2(C.int(len(values)),
		*(**C.VALUE)(unsafe.Pointer(&values))))
}

func EnumFor(self uintptr, values ...uintptr) uintptr {
	if C.rb_block_given_p() != C.int(0) {
		return uintptr(C.Qnil)
	}
	return uintptr(C.rb_funcallv(C.VALUE(self), idEnumFor,
		C.int(len(values)), *(**C.VALUE)(unsafe.Pointer(&values))))
}

func BlockProc() uintptr {
	return uintptr(C.rb_block_proc())
}

func ProcCall(proc uintptr, values ...uintptr) uintptr {
	if len(values) == 0 {
		return uintptr(C.rb_yield(C.Qundef))
	}

	return uintptr(C.rb_funcallv(C.VALUE(proc), idProcCall,
		C.int(len(values)), *(**C.VALUE)(unsafe.Pointer(&values))))
}

func RaiseError(err error) {
	if err == nil {
		return
	}

	strErr := C.VALUE(StringValue(err.Error()))
	C.rb_exc_raise(C.rb_exc_new3(C.rb_eRuntimeError, strErr))
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

func DefineClassInheriting(parent uintptr, name string, super uintptr) uintptr {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	if parent == 0 {
		return uintptr(C.rb_define_class(cname, C.rb_cObject))
	}
	return uintptr(C.rb_define_class_under(C.VALUE(parent),
		cname, C.VALUE(super)))
}

func DefineModule(parent uintptr, name string) uintptr {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	if parent == 0 {
		return uintptr(C.rb_define_module(cname))
	}
	return uintptr(C.rb_define_module_under(C.VALUE(parent), cname))
}
