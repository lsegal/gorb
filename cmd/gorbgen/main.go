package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/lsegal/gorb/codegen"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: gorbgen [options] <package>\n\nOptions:\n")
		flag.PrintDefaults()
		os.Exit(1)
		return
	}

	boolPtr := flag.Bool("build", false, "builds the extension after generating")
	rootPtr := flag.String("root", "", "root Ruby module to place code under")
	flag.Parse()

	if flag.NArg() == 0 {
		flag.Usage()
	}

	gen := codegen.Generator{Path: flag.Arg(0), Root: *rootPtr, Build: *boolPtr}
	gen.Generate()
}
