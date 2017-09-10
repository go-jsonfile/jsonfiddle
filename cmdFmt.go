////////////////////////////////////////////////////////////////////////////
// Program: jsonfiddle
// Purpose: JSON Fiddling
// Authors: Tong Sun (c) 2017, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"bytes"
	"encoding/json"

	"github.com/mkideal/cli"
)

func fmtCLI(ctx *cli.Context) error {
	rootArgv = ctx.RootArgv().(*rootT)
	argv := ctx.Argv().(*fmtT)
	// fmt.Printf("[fmt]:\n  %+v\n  %+v\n  %v\n", rootArgv, argv, ctx.Args())
	Opts.Prefix, Opts.Indent, Opts.Compact, Opts.Protect, Opts.Verbose =
		rootArgv.Prefix, rootArgv.Indent, rootArgv.Compact,
		rootArgv.Protect, rootArgv.Verbose.Value()

	data := readJson(argv.Filei)
	argv.Filei.Close()

	var out bytes.Buffer
	var err error
	if Opts.Compact {
		err = json.Compact(&out, data)
	} else {
		err = json.Indent(&out, data, Opts.Prefix, Opts.Indent)
	}
	abortOn("[::fmt] Formatting input", err)
	out.WriteTo(argv.Fileo)

	return nil
}
