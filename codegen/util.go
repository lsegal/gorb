package codegen

import (
	"fmt"
	"go/ast"
	"path/filepath"
	"regexp"
	"strings"
	"unicode"
)

func resolveType(expr ast.Expr) string {
	switch v := expr.(type) {
	case *ast.ArrayType:
		return "[]" + resolveType(v.Elt)
	case *ast.StarExpr:
		return "*" + resolveType(v.X)
	case *ast.SelectorExpr:
		return v.X.(*ast.Ident).Name + "." + v.Sel.Name
	case *ast.Ident:
		return v.Name
	}

	return "!!UNKNOWN!!"
}

func (g *Generator) resolvedType(s string) string {
	for g.typeAliasMap[s] != "" {
		s = g.typeAliasMap[s]
	}
	return s
}

var reUnderscore = regexp.MustCompile(`([a-z])([A-Z])`)

func underscore(s string) string {
	s = strings.ToLower(reUnderscore.ReplaceAllString(s, "${1}_$2"))
	return s
}

func capitalize(s string) string {
	return strings.ToUpper(s[0:1]) + s[1:]
}

func (g *Generator) returnTypes(t string) ([]string, error) {
	for g.typeAliasMap[t] != "" {
		t = g.typeAliasMap[t]
	}

	switch t {
	case "bool":
		return []string{"BoolValue(bool", "GoBool"}, nil
	case "int":
		return []string{"IntValue(int", "GoInt"}, nil
	case "string":
		return []string{"StringValue(string", "GoString"}, nil
	case "float64":
		return []string{"FloatValue(float64", "GoFloat"}, nil
	}

	if class := g.findClass(t); class != nil {
		return []string{"StructValue", "GoStruct"}, nil
	}

	if strings.Contains(t, ".") {
		pkgName := packageName(t)
		_, ok := g.pkg.imports[pkgName]
		if ok {
			return []string{"ExtStructValue", "GoStruct"}, nil
		}
	}

	return nil, fmt.Errorf("unknown type %s", t)
}

func packageName(t string) string {
	cmps := strings.Split(strings.TrimLeft(t, "*[]"), ".")
	return cmps[0]
}

func isExported(s string) bool {
	s = strings.TrimLeft(s, "*[]")
	if s == "" {
		return false
	}
	return unicode.IsUpper(rune(s[0]))
}

var rePkgReplace = regexp.MustCompile(`^([\*\[\]]*)(\S+)`)

func insertPkg(typ string, pkg string) string {
	return rePkgReplace.ReplaceAllString(typ, "${1}"+pkg+".${2}")
}

func (g *Generator) findClass(name string) *class {
	name = strings.TrimLeft(name, "*[]")
	for _, class := range g.pkg.classes {
		if name == class.name {
			return class
		}
	}
	return nil
}

func (g *Generator) importToModule(name string) string {
	cmps := strings.Split(strings.TrimLeft(name, "[]*"), ".")
	path := g.pkg.imports[cmps[0]]
	r := regexp.MustCompile(g.pkg.importPath + "$")
	base := r.ReplaceAllString(g.pkg.importPackage(), "")
	rel, _ := filepath.Rel(base, path)
	cmps = strings.Split(rel+"/"+cmps[1], "/")
	for i := range cmps {
		cmps[i] = capitalize(cmps[i])
	}
	return strings.Join(cmps, "::")
}

func (g *Generator) writePreambleFunc(name string, arity int) {
	args := make([]string, arity+1)
	for i := range args {
		args[i] = "VALUE"
	}
	fmt.Fprintf(&g.preamble, "extern VALUE %s(%s);\n", name, strings.Join(args, ", "))
}
