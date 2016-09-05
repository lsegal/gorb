package codegen

import (
	"fmt"
	"strings"
	"text/template"
)

type scope int

const (
	instanceScope scope = 0
	classScope          = 1
)

type method struct {
	g           *Generator
	class       *class
	indirection int
	returnClass string
	ctor        bool
	name        string
	scope       scope
	args        []string
	argTypes    []string
	returnType  string
}

func (m *method) ResolvedReturnType() string {
	return m.g.resolvedType(m.returnType)
}

func (m *method) ResolvedReturnClass() string {
	return m.g.resolvedType(m.returnClass)
}

func (m *method) ClassName() string {
	return m.class.Name()
}

func (m *method) ClassVar() string {
	return m.class.VarName()
}

func (m *method) Name() string {
	return capitalize(m.name)
}

func (m *method) RubyName() string {
	name := underscore(m.name)
	if m.ResolvedReturnType() == "bool" {
		name += "?"
	}
	return name
}

func (m *method) FuncName() string {
	s := "i"
	if m.scope == classScope {
		s = "c"
	}
	prefix := "g_" + s + "method_"
	if m.class != nil {
		prefix += m.class.Name()
	}
	return prefix + "_" + m.name
}

func (m *method) Scope() string {
	if m.scope == classScope {
		return "Class"
	}
	return ""
}

func (m *method) RubyArgs() string {
	return strings.Join(append([]string{"self"}, m.args...), ", ")
}

func (m *method) GoArgs() string {
	outargs := make([]string, len(m.args))
	for i, a := range m.args {
		outargs[i] = "go_" + a
	}
	return strings.Join(outargs, ", ")
}

func (m *method) ArgsList() []string {
	return m.args
}

func (m *method) ArgToGo(n int) string {
	return m.typeToGo(m.argTypes[n], m.args[n])
}

func (m *method) ArgToGoExtraArg(n int) string {
	t, _ := m.g.returnTypes(m.argTypes[n])
	if t[1] == "GoStruct" {
		return ", _"
	}
	return ""
}

func (m *method) ReturnTypeToRuby() string {
	ret := "ret"
	if r := m.ResolvedReturnClass(); r != "" && isExported(r) {
		var v string
		if class := m.g.findClass(r); class != nil {
			v = class.VarName()
		} else {
			v = "rb_cObject"
		}
		if m.indirection == 0 {
			ret = "&" + ret
		}
		return fmt.Sprintf("gorb.StructValue(%s, %s)", v, ret)
	}
	t, _ := m.g.returnTypes(m.ResolvedReturnType())
	return fmt.Sprintf("gorb.%s(%s))", t[0], ret)
}

func (m *method) typeToGo(typ string, val string) string {
	t, _ := m.g.returnTypes(typ)
	out := "gorb." + t[1] + "(" + val + ")"
	v := typ
	if isExported(v) {
		for v != "" {
			if m.g.revTypeAliasMap[v] == "" {
				break
			}
			v = m.g.revTypeAliasMap[v]
		}
		v = insertPkg(v, m.g.pkg.name)
	}

	if t[1] == "GoStruct" {
		out = fmt.Sprintf("(%s).(%s)", out, v)
	} else {
		out = fmt.Sprintf("%s(%s)", v, out)
	}
	out = strings.Join(make([]string, m.indirection), "&") + out
	return out
}

func (m *method) ReturnTypeToGo() string {
	return m.typeToGo(m.returnType, "val")
}

func (m *method) ReturnTypeToGoExtraArg() string {
	t, _ := m.g.returnTypes(m.returnType)
	if t[1] == "GoStruct" {
		return ", _"
	}
	return ""
}

func (m *method) HasReturnType() bool {
	return m.returnType != ""
}

func (m *method) FnReceiver() string {
	if m.scope == classScope {
		return m.g.pkg.name
	}
	return "go_obj"
}

const tplMethodData = `
//export {{.FuncName}}
func {{.FuncName}}({{.RubyArgs}} uintptr) uintptr {
{{- if ne .Scope "Class"}}
	{{.FnReceiver}} := g_val2ptr_{{.ClassName}}(self)
{{- end}}
{{- range $i, $v := .ArgsList}}
	go_{{$v}}{{$.ArgToGoExtraArg $i}} := {{$.ArgToGo $i}}
{{- end}} 
{{- if .HasReturnType}}
	ret := {{.FnReceiver}}.{{.Name}}({{.GoArgs}})
	return {{.ReturnTypeToRuby}}
{{- else}}
	{{.FnReceiver}}.{{.Name}}({{.GoArgs}})
	return C.Qnil
{{- end}}
}

`

var tplMethod = template.Must(template.New("method").Parse(tplMethodData))

func (m *method) write(g *Generator) {
	g.writePreambleFunc(m.FuncName(), len(m.args))

	if m.scope == classScope && m.ctor == false { // module function
		fmt.Fprintf(&g.init, `	gorb.DefineModuleFunction(g_pkg, "%s", C.%s, %d)`+"\n",
			m.RubyName(), m.FuncName(), len(m.args))
	} else {
		fmt.Fprintf(&g.init, `	gorb.Define%sMethod(%s, "%s", C.%s, %d)`+"\n",
			m.Scope(), m.class.VarName(), m.RubyName(), m.FuncName(), len(m.args))
	}
	if err := tplMethod.Execute(&g.methods, m); err != nil {
		panic(err)
	}
}
