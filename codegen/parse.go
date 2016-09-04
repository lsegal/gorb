package codegen

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
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
		g.pkg.name = pkg.Name
		g.pkg.importPath = g.Path

		for _, file := range pkg.Files {
			if strings.HasSuffix(file.Name.String(), "_test") {
				continue
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
				g:          g,
				class:      class,
				name:       name.String(),
				returnType: typ,
			}

			if isExported(typ) {
				m.returnClass = typ
			}
			m.indirection = strings.Count(typ, "*")
			m.returnType = typ
			class.attrs = append(class.attrs, &attribute{m})
		}
	}

}

func (g *Generator) parseFunc(f *ast.FuncDecl) {
	if f.Type.Results.NumFields() > 1 {
		fmt.Fprintf(&g.gopreamble, "// skipped func %s() (%s) (unsupported error return)\n",
			f.Name.Name, g.fset.Position(f.Name.Pos()))
		return // TODO support error return
	}

	m := method{g: g, name: f.Name.Name}
	for _, v := range f.Type.Params.List {
		for _, arg := range v.Names {
			m.args = append(m.args, arg.Name)
			m.argTypes = append(m.argTypes, resolveType(v.Type))
		}
	}

	if f.Type.Results.NumFields() == 1 {
		typ := resolveType(f.Type.Results.List[0].Type)
		m.returnType = typ
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
