package codegen

import (
	"fmt"
	"text/template"
)

type class struct {
	pkg      *gpackage
	name     string
	attrs    []*attribute
	imethods []*method
	cmethods []*method
}

func (c *class) VarName() string {
	return "g_class_" + c.name
}

func (c *class) Name() string {
	return capitalize(c.name)
}

func (c *class) FullStructName() string {
	return c.pkg.name + "." + c.name
}

func (c *class) AllocFuncName() string {
	return "g_alloc_" + c.Name()
}

func (c *class) InitFuncName() string {
	return "g_classinit_" + c.Name()
}

const tplClassInitData = `
func g_val2ptr_{{.Name}}(obj uintptr) *{{.FullStructName}} {
	return gorb.GoStruct(obj).(*{{.FullStructName}})
}

//export {{.AllocFuncName}}
func {{.AllocFuncName}}(klass uintptr) uintptr {
	return {{.InitFuncName}}(klass, &{{.FullStructName}}{})
}

func {{.InitFuncName}}(klass uintptr, obj *{{.FullStructName}}) uintptr {
	return gorb.StructValue(klass, obj)
}

`

var tplClassInit = template.Must(template.New("classInit").Parse(tplClassInitData))

func (c *class) write(g *Generator) {
	fmt.Fprintf(&g.gopreamble, "var %s uintptr\n", c.VarName())
	fmt.Fprintf(&g.init, `	%s = gorb.DefineClass(g_pkg, "%s")`+"\n", c.VarName(), c.Name())
	fmt.Fprintf(&g.init, `	gorb.DefineAllocator(%s, C.%s)`+"\n", c.VarName(), c.AllocFuncName())
	g.writePreambleFunc(c.AllocFuncName(), 0)
	if err := tplClassInit.Execute(&g.methods, c); err != nil {
		panic(err)
	}

	for _, attr := range c.attrs {
		attr.write(g)
	}

	for _, m := range c.cmethods {
		m.write(g)
	}

	for _, m := range c.imethods {
		m.write(g)
	}
}
