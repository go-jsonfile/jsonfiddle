////////////////////////////////////////////////////////////////////////////
// Program: jsonfiddle
// Purpose: JSON Fiddling
// Authors: Tong Sun (c) 2017, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"path/filepath"

	"github.com/go-jsonfile/gojson"
	"github.com/mkideal/cli"
	"github.com/suntong/enum"
)

////////////////////////////////////////////////////////////////////////////
// Global variables definitions

var (
	fmtTypeEn enum.Enum
	fmtJson   = fmtTypeEn.Iota("json")
	fmtYaml   = fmtTypeEn.Iota("yaml")

	parser = []gojson.Parser{gojson.ParseJson, gojson.ParseYaml}
)

////////////////////////////////////////////////////////////////////////////
// Function definitions

func j2sCLI(ctx *cli.Context) error {
	rootArgv = ctx.RootArgv().(*rootT)
	argv := ctx.Argv().(*j2sT)
	// fmt.Printf("[j2s]:\n  %+v\n  %+v\n  %v\n", rootArgv, argv, ctx.Args())

	fmtType, ok := fmtTypeEn.Get(argv.FmtType)
	if !ok {
		abortOn("[::j2s]", fmt.Errorf("Invalid format string: '%s'\n", argv.FmtType))
	}
	nameRoot := Basename(filepath.Base(argv.Filei.Name()))
	if len(argv.Name) != 0 {
		nameRoot = argv.Name
	}

	output, err := gojson.Generate(argv.Filei, parser[fmtType],
		nameRoot, argv.Pkg, []string{argv.FmtType}, argv.SubStruct)
	abortOn("[::j2s] Parsing", err)
	fmt.Fprint(argv.Fileo, string(output))
	//abortOn("[::j2s] Writing output", err)

	return nil
}

// func cmdJ2s() error {
// 	return nil
// }
