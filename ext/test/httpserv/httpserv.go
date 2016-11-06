
package main

/*
#include "ruby.h"
extern VALUE g_cmethod__Serve(VALUE, VALUE);

*/
import "C"
import "unsafe"
import "github.com/lsegal/gorb"
import "github.com/lsegal/gorb/test/httpserv"

var _ unsafe.Pointer // ignore unused import warning



//export g_cmethod__Serve
func g_cmethod__Serve(self, addr uintptr) uintptr {
	go_addr := string(gorb.GoString(addr))
	httpserv.Serve(go_addr, block__g_cmethod__Serve)
	return C.Qnil
}

func block__g_cmethod__Serve(arg0 string) (string) {
	rb_arg0 := gorb.StringValue(string(arg0))
	ret := gorb.Yield(rb_arg0)
	return string(gorb.GoString(ret))
}



//export Init_httpserv
func Init_httpserv() {
	g_pkg := gorb.DefineModule(gorb.ModuleRoot, "Test")
	g_pkg = gorb.DefineModule(g_pkg, "Httpserv")

	gorb.DefineModuleFunction(g_pkg, "serve", C.g_cmethod__Serve, 1)

}

func main() {}
