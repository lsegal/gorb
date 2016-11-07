package codegen

import (
	"go/ast"
	"os"
	"path/filepath"
	"strings"
)

type gpackage struct {
	name        string
	importPath  string
	imports     map[string]string
	usedImports map[string]bool
	classes     []*class
	funcs       []*method
	ast         *ast.Package
	noMain      bool
	modNames    []string
}

func (g *gpackage) importPackage() string {
	abs, err := filepath.Abs(g.importPath)
	if err != nil {
		panic(err)
	}

	gopaths := strings.Split(os.Getenv("GOPATH"), ":")
	for _, gopath := range gopaths {
		if !filepath.HasPrefix(abs, gopath) {
			continue
		}

		path, err := filepath.Rel(gopath+"/src", abs)
		if err != nil {
			panic(err)
		}
		return path
	}

	panic("could not find package " + g.importPath + " in GOPATH")
}

func (g *gpackage) allImports() string {
	var out []string
	for pkg, _ := range g.usedImports {
		out = append(out, "import \""+pkg+`"`)
	}
	return strings.Join(out, "\n")
}

func (g *gpackage) moduleName() string {
	return capitalize(g.name)
}
