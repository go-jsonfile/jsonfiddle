////////////////////////////////////////////////////////////////////////////
// Program: jsonfiddle
// Purpose: JSON Fiddling
// Authors: Tong Sun (c) 2017, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/mkideal/cli"
)

func escCLI(ctx *cli.Context) error {
	rootArgv = ctx.RootArgv().(*rootT)
	argv := ctx.Argv().(*escT)
	// fmt.Printf("[esc]:\n  %+v\n  %+v\n  %v\n", rootArgv, argv, ctx.Args())
	Opts.Prefix, Opts.Indent, Opts.Compact, Opts.Verbose =
		rootArgv.Prefix, rootArgv.Indent, rootArgv.Compact, rootArgv.Verbose.Value()

	data, err := ioutil.ReadAll(argv.Filei)
	argv.Filei.Close()
	abortOn("[::esc] Reading input", err)
	// fmt.Printf("] %s", data)

	out, err := json.Marshal(string(data))
	abortOn("[::esc] Formatting input", err)
	fmt.Fprintf(argv.Fileo, "%s", out)

	return nil
}
