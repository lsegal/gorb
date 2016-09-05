
package main

/*
#include "ruby.h"
extern VALUE g_alloc_Node(VALUE);
extern VALUE g_imethod_Node_Value(VALUE);
extern VALUE g_imethod_Node_Value__set(VALUE, VALUE);
extern VALUE g_imethod_Node_Next(VALUE);
extern VALUE g_imethod_Node_Next__set(VALUE, VALUE);
extern VALUE g_cmethod_Node_New(VALUE, VALUE, VALUE);
extern VALUE g_imethod_Node_End(VALUE);

*/
import "C"
import "unsafe"
import "github.com/lsegal/gorb"
import "github.com/lsegal/gorb/test/node"

var _ unsafe.Pointer // ignore unused import warning

var g_class_Node uintptr


func g_val2ptr_Node(obj uintptr) *node.Node {
	return (*node.Node)(gorb.GoStruct(obj))
}

//export g_alloc_Node
func g_alloc_Node(klass uintptr) uintptr {
	return g_classinit_Node(klass, &node.Node{})
}

func g_classinit_Node(klass uintptr, obj *node.Node) uintptr {
	return gorb.StructValue(klass, unsafe.Pointer(obj))
}


//export g_imethod_Node_Value
func g_imethod_Node_Value(self uintptr) uintptr {
	obj := g_val2ptr_Node(self)
	return gorb.StringValue(string(obj.Value))
}

//export g_imethod_Node_Value__set
func g_imethod_Node_Value__set(self, val uintptr) uintptr {
	obj := g_val2ptr_Node(self)
	obj.Value = node.Data(gorb.GoString(val))
	return val
}


//export g_imethod_Node_Next
func g_imethod_Node_Next(self uintptr) uintptr {
	obj := g_val2ptr_Node(self)
	return gorb.StructValue(g_class_Node, unsafe.Pointer(obj.Next))
}

//export g_imethod_Node_Next__set
func g_imethod_Node_Next__set(self, val uintptr) uintptr {
	obj := g_val2ptr_Node(self)
	obj.Next = (*node.Node)(gorb.GoStruct(val))
	return val
}


//export g_cmethod_Node_New
func g_cmethod_Node_New(self, v, n uintptr) uintptr {
	go_v := node.Data(gorb.GoString(v))
	go_n := (*node.Node)(gorb.GoStruct(n))
	ret := node.New(go_v, go_n)
	return gorb.StructValue(g_class_Node, unsafe.Pointer(ret))
}


//export g_imethod_Node_End
func g_imethod_Node_End(self uintptr) uintptr {
	go_obj := g_val2ptr_Node(self)
	ret := go_obj.End()
	return gorb.BoolValue(bool(ret))
}



//export Init_node
func Init_node() {
	g_pkg := gorb.DefineModule(gorb.ModuleRoot, "Test")
	g_pkg = gorb.DefineModule(g_pkg, "Node")

	g_class_Node = gorb.DefineClass(g_pkg, "Node")
	gorb.DefineAllocator(g_class_Node, C.g_alloc_Node)
	gorb.DefineMethod(g_class_Node, "value", C.g_imethod_Node_Value, 0)
	gorb.DefineMethod(g_class_Node, "value=", C.g_imethod_Node_Value__set, 1)
	gorb.DefineMethod(g_class_Node, "next", C.g_imethod_Node_Next, 0)
	gorb.DefineMethod(g_class_Node, "next=", C.g_imethod_Node_Next__set, 1)
	gorb.DefineClassMethod(g_class_Node, "new", C.g_cmethod_Node_New, 2)
	gorb.DefineMethod(g_class_Node, "end?", C.g_imethod_Node_End, 0)

}

func main() {}
