package main

/*
#include "ruby.h"
extern VALUE g_cmethod__IsPrime(VALUE, VALUE);
extern VALUE g_alloc_Fibonacci(VALUE);
extern VALUE g_imethod_Fibonacci_Fib(VALUE, VALUE);

*/
import "C"
import "unsafe"
import "github.com/lsegal/gorb"
import "github.com/lsegal/gorb/test/fib"

var _ unsafe.Pointer // ignore unused import warning

var g_class_Fibonacci uintptr

//export g_cmethod__IsPrime
func g_cmethod__IsPrime(self, n uintptr) uintptr {
	go_n := int(gorb.GoInt(n))
	ret := fib.IsPrime(go_n)
	return gorb.BoolValue(bool(ret))
}

func g_val2ptr_Fibonacci(obj uintptr) *fib.Fibonacci {
	return (*fib.Fibonacci)(gorb.GoStruct(obj))
}

//export g_alloc_Fibonacci
func g_alloc_Fibonacci(klass uintptr) uintptr {
	return g_classinit_Fibonacci(klass, &fib.Fibonacci{})
}

func g_classinit_Fibonacci(klass uintptr, obj *fib.Fibonacci) uintptr {
	return gorb.StructValue(klass, unsafe.Pointer(obj))
}

//export g_imethod_Fibonacci_Fib
func g_imethod_Fibonacci_Fib(self, n uintptr) uintptr {
	go_obj := g_val2ptr_Fibonacci(self)
	go_n := int(gorb.GoInt(n))
	ret := go_obj.Fib(go_n)
	return gorb.IntValue(int(ret))
}

//export Init_fib
func Init_fib() {
	g_pkg := gorb.DefineModule(gorb.ModuleRoot, "Test")
	g_pkg = gorb.DefineModule(g_pkg, "Fib")

	gorb.DefineModuleFunction(g_pkg, "is_prime?", C.g_cmethod__IsPrime, 1)
	g_class_Fibonacci = gorb.DefineClass(g_pkg, "Fibonacci")
	gorb.DefineAllocator(g_class_Fibonacci, C.g_alloc_Fibonacci)
	gorb.DefineMethod(g_class_Fibonacci, "fib", C.g_imethod_Fibonacci_Fib, 1)

}

func main() {}
