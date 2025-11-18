////////////////////////////////////////////////////////////////////////////
// Program: jsonfiddle
// Purpose: JSON Fiddling
// Authors: Tong Sun (c) 2017-2025, All rights reserved
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
	Filei    string `short:"i" long:"input" description:"the source to get json string from (mandatory)" required:"true"`
	Fileo    string `short:"o" long:"output" description:"the output, default to stdout" default:"-"`
	Concise  bool   `short:"s" long:"concise" description:"Compact the top level array into concise array style"`
	Unescape bool   `short:"u" long:"unescape" description:"Unescape unicode of form \u003c to their literal characters"`
}

var fmtCommand FmtCommand

////////////////////////////////////////////////////////////////////////////
// Function definitions

func init() {
	gfParser.AddCommand("fmt",
		"Format json string",
		`
`,
		&fmtCommand)
}

func (x *FmtCommand) Execute(args []string) error {
	fmt.Fprintf(os.Stderr, "Format json string\n")
	// fmt.Fprintf(os.Stderr, "Copyright (C) 2017-2025, Tong Sun\n\n")
	clis.Setup("jsonfiddle::fmt", Opts.Verbose)
	clis.Verbose(1, "Doing Fmt, with %+v, %+v", Opts, args)
	// fmt.Println(x.Filei, x.Fileo, x.Concise, x.Unescape)
	return x.Exec(args)
}

// // Exec implements the business logic of command `fmt`
// func (x *FmtCommand) Exec(args []string) error {
// 	// err := ...
// 	// clis.WarnOn("fmt::Exec", err)
// 	// or,
// 	// clis.AbortOn("fmt::Exec", err)
// 	return nil
// }
