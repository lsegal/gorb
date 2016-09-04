
package main

/*
#include "ruby.h"
extern VALUE g_cmethod__IsPrime(VALUE, VALUE);
extern VALUE g_alloc_Color(VALUE);
extern VALUE g_imethod_Color_R(VALUE);
extern VALUE g_imethod_Color_R__set(VALUE, VALUE);
extern VALUE g_imethod_Color_G(VALUE);
extern VALUE g_imethod_Color_G__set(VALUE, VALUE);
extern VALUE g_imethod_Color_B(VALUE);
extern VALUE g_imethod_Color_B__set(VALUE, VALUE);
extern VALUE g_imethod_Color_HSV(VALUE);
extern VALUE g_imethod_Color_HSV__set(VALUE, VALUE);
extern VALUE g_cmethod_Color_New(VALUE, VALUE, VALUE, VALUE);
extern VALUE g_alloc_Fibonacci(VALUE);
extern VALUE g_imethod_Fibonacci_Fib(VALUE, VALUE);
extern VALUE g_imethod_Fibonacci_Red(VALUE);

*/
import "C"
import "github.com/lsegal/gorb"
import "github.com/lsegal/gorb/test/fib"

var g_class_Color uintptr
var g_class_Fibonacci uintptr


//export g_cmethod__IsPrime
func g_cmethod__IsPrime(self, n uintptr) uintptr {
  go_n := gorb.GoInt(n)
  ret := fib.IsPrime(go_n)
  return gorb.BoolValue(bool(ret))
}


func g_val2ptr_Color(obj uintptr) *fib.Color {
  return gorb.GoStruct(obj).(*fib.Color)
}

//export g_alloc_Color
func g_alloc_Color(klass uintptr) uintptr {
  return g_classinit_Color(klass, &fib.Color{})
}

func g_classinit_Color(klass uintptr, obj *fib.Color) uintptr {
  return gorb.StructValue(klass, obj)
}


//export g_imethod_Color_R
func g_imethod_Color_R(self uintptr) uintptr {
  obj := g_val2ptr_Color(self)
  return gorb.IntValue(int(obj.R))
}

//export g_imethod_Color_R__set
func g_imethod_Color_R__set(self, val uintptr) uintptr {
  obj := g_val2ptr_Color(self)
  obj.R = int(gorb.GoInt(val))
  return val
}


//export g_imethod_Color_G
func g_imethod_Color_G(self uintptr) uintptr {
  obj := g_val2ptr_Color(self)
  return gorb.IntValue(int(obj.G))
}

//export g_imethod_Color_G__set
func g_imethod_Color_G__set(self, val uintptr) uintptr {
  obj := g_val2ptr_Color(self)
  obj.G = int(gorb.GoInt(val))
  return val
}


//export g_imethod_Color_B
func g_imethod_Color_B(self uintptr) uintptr {
  obj := g_val2ptr_Color(self)
  return gorb.IntValue(int(obj.B))
}

//export g_imethod_Color_B__set
func g_imethod_Color_B__set(self, val uintptr) uintptr {
  obj := g_val2ptr_Color(self)
  obj.B = int(gorb.GoInt(val))
  return val
}


//export g_imethod_Color_HSV
func g_imethod_Color_HSV(self uintptr) uintptr {
  obj := g_val2ptr_Color(self)
  return gorb.StringValue(string(obj.HSV))
}

//export g_imethod_Color_HSV__set
func g_imethod_Color_HSV__set(self, val uintptr) uintptr {
  obj := g_val2ptr_Color(self)
  obj.HSV = string(gorb.GoString(val))
  return val
}


//export g_cmethod_Color_New
func g_cmethod_Color_New(self, r, g, b uintptr) uintptr {
  go_r := gorb.GoInt(r)
  go_g := gorb.GoInt(g)
  go_b := gorb.GoInt(b)
  ret := fib.New(go_r, go_g, go_b)
  return gorb.StructValue(g_class_Color, ret)
}


func g_val2ptr_Fibonacci(obj uintptr) *fib.Fibonacci {
  return gorb.GoStruct(obj).(*fib.Fibonacci)
}

//export g_alloc_Fibonacci
func g_alloc_Fibonacci(klass uintptr) uintptr {
  return g_classinit_Fibonacci(klass, &fib.Fibonacci{})
}

func g_classinit_Fibonacci(klass uintptr, obj *fib.Fibonacci) uintptr {
  return gorb.StructValue(klass, obj)
}


//export g_imethod_Fibonacci_Fib
func g_imethod_Fibonacci_Fib(self, n uintptr) uintptr {
  go_obj := g_val2ptr_Fibonacci(self)
  go_n := gorb.GoInt(n)
  ret := go_obj.Fib(go_n)
  return gorb.IntValue(int(ret))
}


//export g_imethod_Fibonacci_Red
func g_imethod_Fibonacci_Red(self uintptr) uintptr {
  go_obj := g_val2ptr_Fibonacci(self)
  ret := go_obj.Red()
  return gorb.StructValue(g_class_Color, &ret)
}



//export Init_fib
func Init_fib() {
  g_pkg := gorb.DefineModule(gorb.ModuleRoot, "Test")
  g_pkg = gorb.DefineModule(g_pkg, "Fib")

  gorb.DefineModuleFunction(g_pkg, "is_prime?", C.g_cmethod__IsPrime, 1)
  g_class_Color = gorb.DefineClass(g_pkg, "Color")
  gorb.DefineAllocator(g_class_Color, C.g_alloc_Color)
  gorb.DefineMethod(g_class_Color, "r", C.g_imethod_Color_R, 0)
  gorb.DefineMethod(g_class_Color, "r=", C.g_imethod_Color_R__set, 1)
  gorb.DefineMethod(g_class_Color, "g", C.g_imethod_Color_G, 0)
  gorb.DefineMethod(g_class_Color, "g=", C.g_imethod_Color_G__set, 1)
  gorb.DefineMethod(g_class_Color, "b", C.g_imethod_Color_B, 0)
  gorb.DefineMethod(g_class_Color, "b=", C.g_imethod_Color_B__set, 1)
  gorb.DefineMethod(g_class_Color, "hsv", C.g_imethod_Color_HSV, 0)
  gorb.DefineMethod(g_class_Color, "hsv=", C.g_imethod_Color_HSV__set, 1)
  gorb.DefineClassMethod(g_class_Color, "new", C.g_cmethod_Color_New, 3)
  g_class_Fibonacci = gorb.DefineClass(g_pkg, "Fibonacci")
  gorb.DefineAllocator(g_class_Fibonacci, C.g_alloc_Fibonacci)
  gorb.DefineMethod(g_class_Fibonacci, "fib", C.g_imethod_Fibonacci_Fib, 1)
  gorb.DefineMethod(g_class_Fibonacci, "red", C.g_imethod_Fibonacci_Red, 0)

}

func main() {}
