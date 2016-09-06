package codegen

import (
	"fmt"
	"regexp"
	"text/template"
)

type attribute struct {
	*method
}

func (a *attribute) ClassName() string {
	return a.class.Name()
}

func (a *attribute) Name() string {
	return a.name
}

func (a *attribute) ReaderRubyName() string {
	return underscore(a.name)
}

func (a *attribute) WriterRubyName() string {
	return underscore(a.name) + "="
}

func (a *attribute) ReaderFuncName() string {
	return "g_imethod_" + a.class.Name() + "_" + a.name
}

func (a *attribute) WriterFuncName() string {
	return "g_imethod_" + a.class.Name() + "_" + a.name + "__set"
}

var reRetReplace = regexp.MustCompile(`\bret\b`)

func (a *attribute) ReturnTypeToRuby() string {
	str := a.method.ReturnTypeToRuby()
	return reRetReplace.ReplaceAllString(str, "obj."+a.Name())
}

func (a *attribute) Indirect() string {
	if a.g.isValueType(a.returnTypes[0]) {
		return "*"
	}
	return ""
}

const tplAttrData = `
//export {{.ReaderFuncName}}
func {{.ReaderFuncName}}(self uintptr) uintptr {
	obj := g_val2ptr_{{.ClassName}}(self)
	return {{.ReturnTypeToRuby}}
}

//export {{.WriterFuncName}}
func {{.WriterFuncName}}(self, val uintptr) uintptr {
	obj := g_val2ptr_{{.ClassName}}(self)
	obj.{{.Name}} = {{.Indirect}}{{.ReturnTypeToGo}}
	return val
}

`

var tplAttr = template.Must(template.New("attr").Parse(tplAttrData))

func (a *attribute) write(g *Generator) {
	fmt.Fprintf(&g.init, `	gorb.DefineMethod(%s, "%s", C.%s, 0)`+"\n",
		a.class.VarName(), a.ReaderRubyName(), a.ReaderFuncName())
	fmt.Fprintf(&g.init, `	gorb.DefineMethod(%s, "%s", C.%s, 1)`+"\n",
		a.class.VarName(), a.WriterRubyName(), a.WriterFuncName())
	g.writePreambleFunc(a.ReaderFuncName(), 0)
	g.writePreambleFunc(a.WriterFuncName(), 1)
	if err := tplAttr.Execute(&g.methods, a); err != nil {
		panic(err)
	}
}
