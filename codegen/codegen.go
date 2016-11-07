package codegen

import (
	"bytes"
	"fmt"
	"go/token"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type Generator struct {
	Path     string
	RootPath string
	OutFile  string
	Root     string
	Build    bool

	code            string
	init            bytes.Buffer
	methods         bytes.Buffer
	gopreamble      bytes.Buffer
	preamble        bytes.Buffer
	fset            token.FileSet
	pkg             gpackage
	outpath         string
	typeAliasMap    map[string]string
	revTypeAliasMap map[string]string
}

func (g *Generator) Generate() {
	g.typeAliasMap = map[string]string{}
	g.revTypeAliasMap = map[string]string{}
	g.parse()
	g.write()
	g.writePath()
	g.build()
}

func (g *Generator) build() {
	if !g.Build {
		return
	}

	dir, _ := os.Getwd()
	os.Chdir(g.outpath)
	defer os.Chdir(dir)

	cmd := exec.Command("make")
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out
	cmd.Run()
	fmt.Print(out.String())
}

func (g *Generator) writePath() {
	g.outpath = filepath.Join(g.RootPath, g.Path)
	fname := g.pkg.name + ".go"
	outfile := filepath.Join(g.outpath, fname)
	if g.OutFile != "" {
		outfile = g.OutFile
	} else {
		os.MkdirAll(g.outpath, 0775)
	}
	ioutil.WriteFile(outfile, []byte(g.code), 0644)
	if g.pkg.noMain {
		return
	}
	ioutil.WriteFile(filepath.Join(g.outpath, "Makefile"), []byte(`
CGO_CFLAGS = -I $(shell ruby -rrbconfig -e 'puts RbConfig::CONFIG["rubyhdrdir"]') -I $(shell ruby -rrbconfig -e 'puts RbConfig::CONFIG["rubyarchhdrdir"]')
CGO_LDFLAGS = -L $(shell ruby -rrbconfig -e 'puts RbConfig::CONFIG["libdir"]') -l$(shell ruby -rrbconfig -e 'puts RbConfig::CONFIG["RUBY_SO_NAME"]')
EXT = $(shell ruby -rrbconfig -e 'puts RbConfig::CONFIG["DLEXT"]')
export CGO_CFLAGS
export CGO_LDFLAGS

all:
	go build -buildmode=c-shared -o `+g.pkg.name+`.${EXT} .
`), 0644)
}

func (g *Generator) rootModule() string {
	var out bytes.Buffer
	modNames := []string{}
	if g.Root != "" {
		modNames = strings.Split(g.Root, "::")
	}

	var pathNames []string
	if len(g.pkg.modNames) > 0 {
		pathNames = g.pkg.modNames
	} else {
		pathNames = strings.Split(strings.Replace(g.Path, ".", "", -1), "/")
	}
	modNames = append(modNames, pathNames...)

	varname := "gorb.ModuleRoot"
	for i, name := range modNames {
		eq := ":="
		if i > 0 {
			eq = "="
		}

		fmt.Fprintf(&out, `	g_pkg %s gorb.DefineModule(%s, "%s")`+"\n",
			eq, varname, capitalize(name))
		varname = "g_pkg"
	}

	return out.String()
}

func (g *Generator) write() {
	for _, fn := range g.pkg.funcs {
		fn.write(g)
	}

	for _, class := range g.pkg.classes {
		class.write(g)
	}

	mainCode := ""
	pkgName := g.pkg.name
	initCode := "func init() {"
	if !g.pkg.noMain {
		pkgName = "main"
		mainCode = "func main() { }"
		initCode = fmt.Sprintf("//export Init_%s\nfunc Init_%s() {",
			g.pkg.name, g.pkg.name)
	}

	g.code = fmt.Sprintf(`package %s

/*
#include "ruby.h"
%s
*/
import "C"
import "unsafe"
import "github.com/lsegal/gorb"
%s

var _ unsafe.Pointer // ignore unused import warning

%s
%s

%s
%s
%s
}

%s
`,
		pkgName, g.preamble.String(), g.pkg.allImports(), g.gopreamble.String(),
		g.methods.String(), initCode, g.rootModule(), g.init.String(), mainCode)

}
