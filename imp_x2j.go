////////////////////////////////////////////////////////////////////////////
// Program: jsonfiddle
// Purpose: JSON Fiddling
// Authors: Tong Sun (c) 2019, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"io"
	"os"

	xj "github.com/basgys/goxml2json"

	"github.com/mkideal/cli"
	"github.com/mkideal/cli/clis"
)

////////////////////////////////////////////////////////////////////////////
// x2j

func x2jCLI(ctx *cli.Context) error {
	rootArgv = ctx.RootArgv().(*rootT)
	argv := ctx.Argv().(*x2jT)
	clis.Setup(progname, rootArgv.Verbose.Value())
	clis.Verbose(2, "<%s> -\n  %+v\n  %+v\n  %v\n", ctx.Path(), rootArgv, argv, ctx.Args())
	Opts.Compact, Opts.Prefix, Opts.Indent, Opts.Protect, Opts.Verbose =
		rootArgv.Compact, rootArgv.Prefix, rootArgv.Indent, rootArgv.Protect, rootArgv.Verbose.Value()
		// argv.Filei, argv.Fileo,

	return DoX2j(argv.Filei, argv.Fileo)
}

// DoX2j implements the business logic of command `x2j`
func DoX2j(cin io.Reader, cout io.Writer) error {
	fmt.Fprintf(os.Stderr, "%s v%s. x2j - XML to JSON\n", progname, version)
	// fmt.Fprintf(os.Stderr, "Copyright (C) 2019, Tong Sun\n\n")

	json, err := xj.Convert(cin)
	abortOn("[::x2j] Convert from xml to json", err)

	fmt.Fprintln(cout, json.String())
	return nil
}
