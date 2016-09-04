package codegen

import (
	"fmt"
	"regexp"
	"strings"
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

func (a *attribute) ReturnTypeToGo() string {
	t, _ := a.g.returnTypes(a.returnType)
	out := "gorb." + t[1] + "(val)"
	v := a.returnType
	if isExported(v) {
		for v != "" {
			if a.g.revTypeAliasMap[v] == "" {
				break
			}
			v = a.g.revTypeAliasMap[v]
		}
		v = insertPkg(v, a.g.pkg.name)
	}

	if t[1] == "GoStruct" {
		out = fmt.Sprintf("(%s).(%s)", out, v)
	} else {
		out = fmt.Sprintf("%s(%s)", v, out)
	}
	out = strings.Join(make([]string, a.indirection), "&") + out
	return out
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
  obj.{{.Name}} = {{.ReturnTypeToGo}}
  return val
}

`

var tplAttr = template.Must(template.New("attr").Parse(tplAttrData))

func (a *attribute) write(g *Generator) {
	fmt.Fprintf(&g.init, `  gorb.DefineMethod(%s, "%s", C.%s, 0)`+"\n",
		a.class.VarName(), a.ReaderRubyName(), a.ReaderFuncName())
	fmt.Fprintf(&g.init, `  gorb.DefineMethod(%s, "%s", C.%s, 1)`+"\n",
		a.class.VarName(), a.WriterRubyName(), a.WriterFuncName())
	g.writePreambleFunc(a.ReaderFuncName(), 0)
	g.writePreambleFunc(a.WriterFuncName(), 1)
	if err := tplAttr.Execute(&g.methods, a); err != nil {
		panic(err)
	}
}
