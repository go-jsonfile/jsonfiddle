////////////////////////////////////////////////////////////////////////////
// Program: jf
// Purpose: JSON Fiddling
// Authors: Tong Sun (c) 2017, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/mkideal/cli"
)

////////////////////////////////////////////////////////////////////////////
// sort

func sortCLI(ctx *cli.Context) error {
	rootArgv = ctx.RootArgv().(*rootT)
	argv := ctx.Argv().(*sortT)
	//fmt.Printf("[sort]:\n  %+v\n  %+v\n  %v\n", rootArgv, argv, ctx.Args())
	Opts.Prefix, Opts.Indent, Opts.Compact, Opts.Protect, Opts.Verbose =
		rootArgv.Prefix, rootArgv.Indent, rootArgv.Compact,
		rootArgv.Protect, rootArgv.Verbose.Value()
	return cmdSort(argv.Filei, argv.Fileo)
}

func cmdSort(r io.Reader, w io.Writer) error {
	var res interface{}
	content := readJson(r)
	json.Unmarshal(content, &res)
	var js []byte
	var err error
	if Opts.Compact {
		js, err = json.Marshal(res)
	} else {
		js, err = json.MarshalIndent(res, Opts.Prefix, Opts.Indent)
	}
	abortOn("[::sort] Marshaling input", err)
	fmt.Fprintln(w, string(js))
	return nil
}
