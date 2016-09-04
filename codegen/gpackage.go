package codegen

import (
	"os"
	"path/filepath"
	"strings"
)

type gpackage struct {
	name       string
	importPath string
	classes    []*class
	funcs      []*method
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

func (g *gpackage) moduleName() string {
	return capitalize(g.name)
}
