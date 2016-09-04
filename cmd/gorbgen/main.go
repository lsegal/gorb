package main

import (
	"flag"

	"github.com/lsegal/gorb/codegen"
)

func main() {
	boolPtr := flag.Bool("build", false, "builds the extension after generating")
	rootPtr := flag.String("root", "", "root Ruby module to place code under")
	flag.Parse()

	gen := codegen.Generator{Path: flag.Arg(0), Root: *rootPtr, Build: *boolPtr}
	gen.Generate()
}
