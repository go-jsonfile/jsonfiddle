////////////////////////////////////////////////////////////////////////////
// Program: jsonfiddle
// Purpose: JSON Fiddling
// Authors: Tong Sun (c) 2017-2023, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"os"

	"github.com/go-easygen/go-flags/clis"
)

// *** Sub-command: fmt ***

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

// The FmtCommand type defines all the configurable options from cli.
type FmtCommand struct {
	Filei string `short:"i" long:"input" description:"the source to get json string from (mandatory)" required:"true"`
	Fileo string `short:"o" long:"output" description:"the output, default to stdout" default:"-"`
}

var fmtCommand FmtCommand

func init() {
	parser.AddCommand("fmt",
		"Format json string",
		"",
		&fmtCommand)
}

func (x *FmtCommand) Execute(args []string) error {
	fmt.Fprintf(os.Stderr, "Format json string\n")
	// fmt.Fprintf(os.Stderr, "Copyright (C) 2017-2023, Tong Sun\n\n")
	clis.Setup("jsonfiddle::fmt", opts.Verbose)
	clis.Verbose(1, "Doing Fmt, with %+v, %+v", opts, args)
// 	fmt.Println(x.Filei, x.Fileo)
	return x.Exec(args)
}

// Exec implements the business logic of command `fmt`
// func (x *FmtCommand) Exec(args []string) error {
// 	// err := ...
// 	// clis.WarnOn("Fmt, Exec", err)
// 	// or,
// 	// clis.AbortOn("Fmt, Exec", err)
// 	return nil
// }
