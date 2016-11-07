package main

import (
	"flag"
	"fmt"
	"os"
	"path"

	"github.com/lsegal/gorb/codegen"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: gorbgen [options] <package> [package ...]\n\nOptions:\n")
		flag.PrintDefaults()
		os.Exit(1)
		return
	}

	boolPtr := flag.Bool("build", false, "builds the extension after generating")
	rootPtr := flag.String("root", "", "root Ruby module to place code under")
	rootPathPtr := flag.String("root-path", "ext", "root directory on disk to place code")
	outFilePtr := flag.String("file", "", "override the filename to write to")
	flag.Parse()

	if flag.NArg() == 0 {
		flag.Usage()
	}

	for i := 0; i < flag.NArg(); i++ {
		gen := codegen.Generator{
			Path:     path.Clean(flag.Arg(i)),
			Root:     *rootPtr,
			RootPath: *rootPathPtr,
			OutFile:  *outFilePtr,
			Build:    *boolPtr,
		}
		gen.Generate()
	}
}
