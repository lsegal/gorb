package codegen

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"path"
	"strconv"
	"strings"
)

func (g *Generator) parse() {
	pkgs, err := parser.ParseDir(&g.fset, g.Path, nil, 0)
	if err != nil {
		panic(err)
	}

	for _, pkg := range pkgs {
		if strings.HasSuffix(pkg.Name, "_test") {
			continue
		}

		ast.PackageExports(pkg)
		g.pkg.ast = pkg
		g.pkg.imports = map[string]string{}
		g.pkg.name = pkg.Name
		g.pkg.importPath = g.Path
		g.pkg.usedImports = map[string]bool{}
		g.pkg.usedImports[g.pkg.importPackage()] = true

		for _, file := range pkg.Files {
			if strings.HasSuffix(file.Name.String(), "_test") {
				continue
			}

			for _, i := range file.Imports {
				g.addPackage(i)
			}

			for _, decl := range file.Decls {
				switch d := decl.(type) {
				case *ast.GenDecl:
					g.parseType(d)
				case *ast.FuncDecl:
					g.parseFunc(d)
				}
			}
		}

		break
	}
}

func (g *Generator) addPackage(i *ast.ImportSpec) {
	ipath := strings.Replace(i.Path.Value, `"`, "", -1)

	var name string
	if i.Name != nil {
		name = i.Name.String()
	} else {
		name = path.Base(ipath)
	}

	if strings.HasPrefix(ipath, path.Dir(g.pkg.importPackage())) {
		g.pkg.imports[name] = ipath
	}
}

func (g *Generator) parseType(d *ast.GenDecl) {
	if d.Tok != token.TYPE {
		return
	}

	for _, spec := range d.Specs {
		tspec := spec.(*ast.TypeSpec)
		switch s := tspec.Type.(type) {
		case *ast.StructType:
			c := &class{pkg: &g.pkg, name: tspec.Name.String()}
			for _, field := range s.Fields.List {
				g.parseField(field, c)
			}
			g.pkg.classes = append(g.pkg.classes, c)

		default:
			t := resolveType(s)
			g.typeAliasMap[tspec.Name.String()] = t
			g.revTypeAliasMap[t] = tspec.Name.String()
		}
	}
}

func (g *Generator) parseField(field *ast.Field, class *class) {
	for _, name := range field.Names {
		if name.IsExported() {
			typ := resolveType(field.Type)
			m := &method{
				g:           g,
				class:       class,
				name:        name.String(),
				returnTypes: []string{typ},
			}

			if isExported(typ) {
				m.returnClass = typ
			}
			m.indirection = strings.Count(typ, "*")
			class.attrs = append(class.attrs, &attribute{m})
		}
	}

}

func (g *Generator) parseBlockArg(m *method, f *ast.FuncType) {
	for i, v := range f.Params.List {
		names := append([]*ast.Ident{}, v.Names...)
		if len(names) == 0 {
			names = []*ast.Ident{ast.NewIdent("arg" + strconv.Itoa(i))}
		}
		for _, arg := range names {
			m.blockArgs = append(m.blockArgs, arg.Name)
			m.blockArgTypes = append(m.blockArgTypes, resolveType(v.Type))
		}
	}

	for _, rf := range f.Results.List {
		m.blockReturnTypes = append(m.blockReturnTypes, resolveType(rf.Type))
	}
}

func (g *Generator) parseFunc(f *ast.FuncDecl) {
	m := method{g: g, name: f.Name.Name}
	for i, v := range f.Type.Params.List {
		switch t := v.Type.(type) {
		case *ast.FuncType:
			if i == len(f.Type.Params.List)-1 {
				g.parseBlockArg(&m, t)
			}
		default:
			for _, arg := range v.Names {
				m.args = append(m.args, arg.Name)
				m.argTypes = append(m.argTypes, resolveType(v.Type))
			}
		}
	}

	appendReturnTypes(&m, f.Type)
	if f.Type.Results.NumFields() == 1 {
		typ := resolveType(f.Type.Results.List[0].Type)
		if isExported(typ) {
			m.returnClass = typ
		}
		m.indirection = strings.Count(typ, "*")
	}

	if f.Recv.NumFields() > 0 { // instance method
		typ := resolveType(f.Recv.List[0].Type)
		if class := g.findClass(typ); class != nil {
			m.class = class
			m.scope = instanceScope
			class.imethods = append(class.imethods, &m)
			return
		} else if isExported(typ) {
			fmt.Fprintf(&g.gopreamble, "// skipped func %s() (%s) (could not resolve class %s)\n",
				f.Name.Name, g.fset.Position(f.Name.Pos()), typ)
			return
		}
	} else if f.Type.Results.NumFields() == 1 {
		typ := resolveType(f.Type.Results.List[0].Type)
		if class := g.findClass(typ); class != nil { // ctor
			m.class = class
			m.scope = classScope
			m.ctor = true
			m.returnClass = class.name
			class.cmethods = append(class.cmethods, &m)
			return
		} else if isExported(typ) {
			fmt.Fprintf(&g.gopreamble, "// skipped func %s() (%s) (could not resolve class %s)\n",
				f.Name.Name, g.fset.Position(f.Name.Pos()), typ)
			return
		}
	}

	m.scope = classScope
	g.pkg.funcs = append(g.pkg.funcs, &m)
}

func appendReturnTypes(m *method, t *ast.FuncType) {
	for _, rf := range t.Results.List {
		m.returnTypes = append(m.returnTypes, resolveType(rf.Type))
	}
}
