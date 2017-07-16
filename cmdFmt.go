////////////////////////////////////////////////////////////////////////////
// Program: jsonfiddle
// Purpose: JSON Fiddling
// Authors: Tong Sun (c) 2017, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"

	"github.com/mkideal/cli"
)

func fmtCLI(ctx *cli.Context) error {
	rootArgv = ctx.RootArgv().(*rootT)
	argv := ctx.Argv().(*fmtT)
	// fmt.Printf("[fmt]:\n  %+v\n  %+v\n  %v\n", rootArgv, argv, ctx.Args())
	Opts.Prefix, Opts.Indent, Opts.Compact, Opts.Verbose =
		rootArgv.Prefix, rootArgv.Indent, rootArgv.Compact, rootArgv.Verbose.Value()

	data, err := ioutil.ReadAll(argv.Filei)
	argv.Filei.Close()
	abortOn("[::fmt] Reading input", err)

	var out bytes.Buffer
	if Opts.Compact {
		err = json.Compact(&out, data)
	} else {
		err = json.Indent(&out, data, Opts.Prefix, Opts.Indent)
	}
	abortOn("[::fmt] Formatting input", err)
	out.WriteTo(argv.Fileo)

	return nil
}
